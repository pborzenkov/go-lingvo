package lingvo

import (
	"net/http"
	"net/url"
)

// config contains tunables for Lingvo API client
type config struct {
	httpClient *http.Client
	baseURL    *url.URL
	userAgent  string
}

var defaultConfig = config{
	httpClient: http.DefaultClient,
	baseURL:    mustParseURL(defaultBaseURL),
	userAgent:  defaultUserAgent,
}

// Option configure Lingvo API client
type Option func(cfg *config)

// WithHTTPClient sets custom HTTP client for the API client. If not set,
// http.DefaultClient is used.
func WithHTTPClient(client *http.Client) Option {
	return func(cfg *config) {
		if client == nil {
			panic("HTTP client can't be nil")
		}
		cfg.httpClient = client
	}
}

// WithBaseURL sets custom API base URL.
func WithBaseURL(baseURL *url.URL) Option {
	return func(cfg *config) {
		if baseURL == nil {
			panic("base URL can't be nil")
		}
		cfg.baseURL = baseURL
	}
}

// WithUserAgent sets custom user agent for the API client.
func WithUserAgent(userAgent string) Option {
	return func(cfg *config) {
		cfg.userAgent = userAgent
	}
}

func mustParseURL(urlStr string) *url.URL {
	u, err := url.Parse(urlStr)
	if err != nil {
		panic(err)
	}

	return u
}
