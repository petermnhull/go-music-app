package endpoints

type apiResponse struct {
	Code   int64       `json:"code"`
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}
