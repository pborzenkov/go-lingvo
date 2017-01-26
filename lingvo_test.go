package lingvo

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
	"strings"
	"testing"
)

var (
	mux    *http.ServeMux
	server *httptest.Server

	apiKey string
	client *Client
)

func setup() {
	mux = http.NewServeMux()
	server = httptest.NewServer(mux)
	url, _ := url.Parse(server.URL)

	apiKey = "api_key"
	client = NewClient(apiKey, WithBaseURL(url))
}

func teardown() {
	server.Close()
}

func testMethod(t *testing.T, r *http.Request, want string) {
	if got := r.Method; got != want {
		t.Errorf("unexpected request method, want = %q, got = %q", want, got)
	}
}

type values map[string]string

func testFormValues(t *testing.T, r *http.Request, values values) {
	want := url.Values{}
	for k, v := range values {
		want.Add(k, v)
	}

	r.ParseForm()
	if got := r.Form; !reflect.DeepEqual(got, want) {
		t.Errorf("unexpected request parameters, want = %q, got = %q", want, got)
	}
}

func TestNewClient(t *testing.T) {
	c := NewClient("")

	if got, want := c.cfg.baseURL.String(), defaultBaseURL; got != want {
		t.Errorf("NewClient BaseURL is %v, want %v", got, want)
	}
	if got, want := c.cfg.userAgent, defaultUserAgent; got != want {
		t.Errorf("NewClient UserAgent is %v, want %v", got, want)
	}
}

func TestNewRequest(t *testing.T) {
	c := NewClient("")

	inURL, outURL := "endpoint", defaultBaseURL+"endpoint"
	req, _ := c.NewRequest("GET", inURL, nil)

	if got, want := req.URL.String(), outURL; got != want {
		t.Errorf("NewRequest(%q) URL is %v, want %v", inURL, got, want)
	}

	if got, want := req.Header.Get("User-Agent"), c.cfg.userAgent; got != want {
		t.Errorf("NewRequest() User-Agent is %v, want %v", got, want)
	}
}

func TestDo(t *testing.T) {
	setup()
	defer teardown()

	type foo struct {
		A string
		B string
	}

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")

		fmt.Fprint(w, `{"A":"a","B":"b"}`)
	})

	req, _ := client.NewRequest("GET", "/", nil)
	body := new(foo)
	client.Do(context.Background(), req, body)

	want := &foo{"a", "b"}
	if !reflect.DeepEqual(body, want) {
		t.Errorf("Response body = %v, want %v", body, want)
	}
}

func TestDo_httpError(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "some error", 400)
	})

	req, _ := client.NewRequest("GET", "/", nil)
	err := client.Do(context.Background(), req, nil)

	if err == nil {
		t.Error("Expected HTTP 400 error.")
	}

	if _, ok := err.(*ErrorResponse); !ok {
		t.Errorf("Expected a *ErrorResponse error; got %#v.", err)
	}
}

func TestDo_authorization(t *testing.T) {
	setup()
	defer teardown()

	type foo struct {
		A string
		B string
	}

	token := "token"
	authorized := false
	mux.HandleFunc("/"+endpointAuth, func(w http.ResponseWriter, r *http.Request) {
		if got, want := r.Header.Get("Authorization"), "Basic "+apiKey; got != want {
			t.Errorf("Authorization header = %v, want %v", got, want)
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		fmt.Fprintf(w, "%s", token)
		authorized = true
	})
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if got, want := r.Header.Get("Authorization"), "Bearer "+token; got != want {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			if authorized {
				t.Errorf("Authorization header = %v, want %v", got, want)
			}
			return
		}
		fmt.Fprint(w, `{"A":"a","B":"b"}`)
	})

	req, _ := client.NewRequest("GET", "/", nil)
	body := new(foo)
	err := client.Do(context.Background(), req, body)
	if err != nil {
		t.Errorf("Unexpected error response \"%v\"", err)
	}
	want := &foo{"a", "b"}
	if !reflect.DeepEqual(body, want) {
		t.Errorf("Response body = %v, want %v", body, want)
	}
}

func TestNewRequest_badURL(t *testing.T) {
	c := NewClient("")
	_, err := c.NewRequest("GET", ":", nil)

	if err == nil {
		t.Errorf("Expected error to be returned")
	}
	if err, ok := err.(*url.Error); !ok || err.Op != "parse" {
		t.Errorf("Expected URL parse error, got %+v", err)
	}
}

func TestCheckResponse(t *testing.T) {
	c := NewClient("")

	res := &http.Response{
		Request:    &http.Request{},
		StatusCode: http.StatusBadRequest,
		Body:       ioutil.NopCloser(strings.NewReader("some error")),
	}
	err := c.checkResponse(res)
	if err == nil {
		t.Errorf("Expected error response")
	}

	rerr := err.(*ErrorResponse)
	want := &ErrorResponse{
		Response: res,
		Message:  "some error",
	}
	if !reflect.DeepEqual(rerr, want) {
		t.Errorf("Error = %#v, want %#v", err, want)
	}
}

func TestCheckResponse_noBody(t *testing.T) {
	c := NewClient("")

	res := &http.Response{
		Request:    &http.Request{},
		StatusCode: http.StatusBadRequest,
		Body:       ioutil.NopCloser(strings.NewReader("")),
	}
	err := c.checkResponse(res)
	if err == nil {
		t.Errorf("Expected error response")
	}

	rerr := err.(*ErrorResponse)
	want := &ErrorResponse{
		Response: res,
	}
	if !reflect.DeepEqual(rerr, want) {
		t.Errorf("Error = %#v, want %#v", err, want)
	}
}

func TestCheckResponse_unauthorized(t *testing.T) {
	c := NewClient("token")

	res := &http.Response{
		Request:    &http.Request{},
		StatusCode: http.StatusUnauthorized,
		Body:       ioutil.NopCloser(strings.NewReader("unauthorized")),
	}

	err := c.checkResponse(res)
	if err == nil {
		t.Errorf("Expected error response")
	}

	if _, ok := err.(*UnauthorizedError); !ok {
		t.Errorf("Expected UnauthorizedError error")
	}
}
