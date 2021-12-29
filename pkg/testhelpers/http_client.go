package testhelpers

import (
	"bytes"
	"io"
	"net/http"
)

// MockHTTPClient provides basic mock http client
type MockHTTPClient struct {
	Responses map[string]http.Response
}

// NewMockHTTPClient initialises a mock
func NewMockHTTPClient() *MockHTTPClient {
	responses := make(map[string]http.Response)
	return &MockHTTPClient{responses}
}

// AddMockResponse creates a mapping from URL to expected response
func (m *MockHTTPClient) AddMockResponse(url string, statusCode int, data string) {
	response := http.Response{StatusCode: statusCode, Body: io.NopCloser(bytes.NewBuffer([]byte(data)))}
	m.Responses[url] = response
}

// Do for satisfying http client interface
func (m MockHTTPClient) Do(r *http.Request) (*http.Response, error) {
	resp, ok := m.Responses[r.URL.String()]
	if !ok {
		return &http.Response{StatusCode: http.StatusNotFound}, nil
	}
	return &resp, nil
}
