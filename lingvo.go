package lingvo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"sync"

	"golang.org/x/net/context"
)

const (
	libraryVersion   = "0.0.0"
	defaultBaseURL   = "https://developers.lingvolive.com/"
	defaultUserAgent = "go-lingvo/" + libraryVersion

	endpointAuth = "api/v1.1/authenticate"
)

// Client manages communication with Lingvo API. It is safe to use it from
// multiple goroutines.
type Client struct {
	cfg config

	apiKey string

	mu    sync.RWMutex // protects token
	token string
}

// NewClient returns a new Lingvo API client customized using the provided
// slice of options.
// The client authenticates with the API server using the provided apiKey.
func NewClient(apiKey string, opts ...Option) *Client {
	cl := &Client{
		cfg:    defaultConfig,
		apiKey: apiKey,
	}

	for _, opt := range opts {
		opt(&cl.cfg)
	}

	return cl
}

// NewRequest creates an API request. A relative URL can be provided in urlStr,
// in which case it is resolved relative to the BaseURL of the Client. body is
// currently ignored.
func (c *Client) NewRequest(method, urlStr string, body interface{}) (*http.Request, error) {
	return c.newRequest(method, urlStr, body, nil)
}

func (c *Client) newRequest(method, urlStr string, body interface{}, headers map[string]string) (*http.Request, error) {
	rel, err := url.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	u := c.cfg.baseURL.ResolveReference(rel)

	req, err := http.NewRequest(method, u.String(), nil)
	if err != nil {
		return nil, err
	}

	if c.cfg.userAgent != "" {
		req.Header.Set("User-Agent", c.cfg.userAgent)
	}
	for k, v := range headers {
		req.Header.Set(k, v)
	}

	return req, nil
}

// Do sends an API request and returns the API response.  The API response is
// JSON decoded and stored in the value pointed to by v.  If v implements the
// io.Writer interface, the raw response body will be written to v, without
// attempting to first decode it.
func (c *Client) Do(ctx context.Context, req *http.Request, v interface{}) error {
	return doWithRetry(ctx, func() error {
		// do this on every iteration, because new token may be
		// received after retry
		c.mu.RLock()
		if c.token != "" {
			req.Header.Set("Authorization", "Bearer "+c.token)
		}
		c.mu.RUnlock()

		return c.do(ctx, req, v)
	})
}

func (c *Client) do(ctx context.Context, req *http.Request, v interface{}) error {
	resp, err := c.cfg.httpClient.Do(req.WithContext(ctx))
	if err != nil {
		return err
	}

	defer func() {
		io.CopyN(ioutil.Discard, resp.Body, 512)
		resp.Body.Close()
	}()

	if err = c.checkResponse(resp); err != nil {
		return err
	}

	if v != nil {
		if w, ok := v.(io.Writer); ok {
			_, err = io.Copy(w, resp.Body)
		} else {
			err = json.NewDecoder(resp.Body).Decode(v)
		}
	}

	return err
}

func (c *Client) checkResponse(r *http.Response) error {
	if r.StatusCode >= 200 && r.StatusCode <= 299 {
		return nil
	}

	msg := ""
	data, err := ioutil.ReadAll(r.Body)
	if err == nil && data != nil {
		msg = string(data)
	}

	switch r.StatusCode {
	case http.StatusUnauthorized:
		// TODO: While undocumented, the token is actually base64
		// encoded JWT.  Consider looking at its expiration time instead
		// of relying purely on 401.
		// TODO: Consider updating token in a dedicated goroutine, so
		// that multiple simultaneous StatusUnauthorized responses from
		// several goroutines don't trigger multiple token refresh
		// requests.
		return &UnauthorizedError{
			Response:    r,
			Message:     msg,
			maxAttempts: 2,
			preRetry: func(ctx context.Context) error {
				var token bytes.Buffer

				if c.apiKey == "" {
					return fmt.Errorf("API key is missing")
				}
				req, err := c.newRequest("POST", endpointAuth, nil, map[string]string{
					"Authorization": "Basic " + c.apiKey,
				})
				if err != nil {
					return err
				}
				if err = c.do(ctx, req, &token); err != nil {
					return err
				}
				c.mu.Lock()
				c.token = token.String()
				c.mu.Unlock()
				return nil
			},
		}
	default:
		return &ErrorResponse{
			Response: r,
			Message:  msg,
		}
	}
}

func doWithRetry(ctx context.Context, fn func() error) error {
	var err error
	var attempts int

	for {
		attempts++
		err = fn()
		rerr, ok := err.(RetryableError)
		if !ok || attempts >= rerr.MaxAttempts() {
			return err
		}
		if err = rerr.PreRetry()(ctx); err != nil {
			return err
		}
	}
}

// RetryableError is an error that indicates that an action that caused the
// error should be retried.
type RetryableError interface {
	Error() string
	// MaxAttempts returns maximum number of attempts to retry the error.
	MaxAttempts() int
	// PreRetry return a function that should be executed before trying to
	// retry the action.
	PreRetry() func(ctx context.Context) error
}

// ErrorResponse reports an error caused by an API request.
type ErrorResponse struct {
	Response *http.Response // HTTP response that caused this error
	Message  string         // error message
}

// Error implements error interface
func (r *ErrorResponse) Error() string {
	return fmt.Sprintf("%v %v: %d %v", r.Response.Request.Method, r.Response.Request.URL,
		r.Response.StatusCode, r.Message)
}

// UnauthorizedError occurs when Lingvo API returns 401 "Unauthorized"
type UnauthorizedError struct {
	Response *http.Response // HTTP response that caused this error
	Message  string         // error message

	maxAttempts int                             // number of attempts to retry the request that caused this error
	preRetry    func(ctx context.Context) error // function to call before retrying the request
}

// Error implements RetryableError interface
func (r *UnauthorizedError) Error() string {
	return fmt.Sprintf("%v %v: %d %v", r.Response.Request.Method, r.Response.Request.URL,
		r.Response.StatusCode, r.Message)
}

// MaxAttempts implements RetryableError interface
func (r *UnauthorizedError) MaxAttempts() int {
	return r.maxAttempts
}

// PreRetry implements RetryableError interface
func (r *UnauthorizedError) PreRetry() func(ctx context.Context) error {
	return r.preRetry
}
