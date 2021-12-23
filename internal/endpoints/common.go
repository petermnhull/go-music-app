package endpoints

type APIResponseData interface{}

type APIResponse struct {
	Code   int64           `json:"code"`
	Status string          `json:"status"`
	Data   APIResponseData `json:"data"`
}

func NewFailedAPIResponse(code int64, message string) *APIResponse {
	data := map[string]string{
		"error": message,
	}
	response := APIResponse{code, "failed", data}
	return &response
}
