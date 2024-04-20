package response

type DefaultResponse struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func CreateResponse(status string, message string, data interface{}) (response DefaultResponse) {
	response = DefaultResponse{
		Status:  status,
		Message: message,
		Data:    data,
	}

	return
}

func ErrorResponse(status, message string) (response DefaultResponse) {
	response = DefaultResponse{
		Status:  status,
		Message: message,
		Data:    nil,
	}
	return
}
