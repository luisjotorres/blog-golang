package domain

import "fmt"

type APIError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Status  string `json:"status"`
}

func NewAPIError(code int, message, status string) *APIError {
	return &APIError{
		code,
		message,
		status,
	}
}

func (ae *APIError) Error() string {
	return fmt.Sprintf("%s - %s", ae.Message, ae.Status)
}
