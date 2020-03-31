package errors

import "net/http"

type RestErr struct {
	Success bool   `json:"success" default0:"false"`
	Message string `json:"message"`
	Status  int    `json:"status"`
}

func NewBadRequestErr(message string) *RestErr {
	return &RestErr{
		Message: message,
		Status:  http.StatusBadRequest,
	}
}

func NewNotFoundRequestErr(message string) *RestErr {
	return &RestErr{
		Message: message,
		Status:  http.StatusNotFound,
	}
}
func NewInternalServerError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Status:  http.StatusInternalServerError,
	}
}

func NewUnauthorizedError() *RestErr {
	return &RestErr{
		Message: "Unauthorized",
		Status:  http.StatusUnauthorized,
	}
}
