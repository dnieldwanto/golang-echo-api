package utils

type WebResponse struct {
	Status  int         `json:"status"`
	Message string      `json:"Message"`
	Data    interface{} `json:"data"`
}

func GenerateResponse(status int, message string, data interface{}) WebResponse {
	return WebResponse{
		Status:  status,
		Message: message,
		Data:    data,
	}
}

func GenerateResponseV2(status int, data interface{}) WebResponse {
	return WebResponse{
		Status:  status,
		Message: "OK",
		Data:    data,
	}
}
