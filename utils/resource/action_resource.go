package resource

type request struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func ActionResource(success bool, message string, data ...interface{}) *request {
	return &request{
		Success: success,
		Message: message,
		Data:    data,
	}
}
