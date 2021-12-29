package services

import "net/http"

// HTTPClientInterface for making requests
type HTTPClientInterface interface {
	Do(request *http.Request) (*http.Response, error)
}
