package pkg

import "net/http"

// HTTPClient for making requests
type HTTPClient interface {
	Do(request *http.Request) (*http.Response, error)
}
