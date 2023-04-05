package httputil

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
)

// Modifier modifies the Request before it's sent out to add extra info.
type Modifier func(c *Client, req *http.Request)

// WithHeader attaches the provided headers to the request.
func WithHeader(header map[string][]string) Modifier {
	return func(_ *Client, req *http.Request) {
		for k, v := range header {
			for _, w := range v {
				req.Header.Add(k, w)
			}
		}
	}
}

// WithToken attaches the AccessToken to the request.
// It should be included with requests that require authentication.
func WithToken() Modifier {
	return func(c *Client, req *http.Request) {
		req.Header.Add("Authorization", "Bearer "+c.AccessToken)
	}
}

// WithBody attaches a io.ReadCloser to the request.
func WithBody(body io.ReadCloser) Modifier {
	return func(_ *Client, req *http.Request) {
		req.Body = body
	}
}

// WithJSONBody attaches a JSON body to the request.
func WithJSONBody(body interface{}) Modifier {
	return func(_ *Client, req *http.Request) {
		b, err := json.Marshal(body)
		if err != nil {
			return
		}

		req.Body = io.NopCloser(bytes.NewReader(b))
		req.ContentLength = int64(len(b))
		req.Header.Add("Content-Type", "application/json")
	}
}

// WithQuery attaches one-to-one queries to the request.
// It is provided as a helper function that calls WithFullQuery.
func WithQuery(rawQueries map[string]string) Modifier {
	fullQuery := make(map[string][]string)
	for k, v := range rawQueries {
		fullQuery[k] = []string{v}
	}
	return WithFullQuery(fullQuery)
}

// WithFullQuery attaches one-to-many queries to the request.
func WithFullQuery(query map[string][]string) Modifier {
	encoded := url.Values(query).Encode()
	return func(_ *Client, req *http.Request) {
		req.URL.RawQuery = encoded
	}
}
