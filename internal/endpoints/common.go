package endpoints

import "encoding/json"

// APIResponseData holds generic data for API response
type APIResponseData interface{}

// APIResponse holds general information to return
type APIResponse struct {
	Code   int64           `json:"code"`
	Status string          `json:"status"`
	Data   APIResponseData `json:"data"`
}

// ToOutput converts APIResponse into JSON string output
func (r *APIResponse) ToOutput() string {
	responseBytes, _ := json.Marshal(r)
	output := string(responseBytes)
	return output
}

// NewAPIResponseSuccess creates a new APIResponse for a successful request
func NewAPIResponseSuccess(code int64, data interface{}) *APIResponse {
	response := APIResponse{code, "success", data}
	return &response
}

// NewAPIResponseFailed creates a new APIResponse for a failed request
func NewAPIResponseFailed(code int64, message string) *APIResponse {
	data := map[string]string{"error": message}
	response := APIResponse{code, "failed", data}
	return &response
}
