package utils

type ErrorWithJSONFormat struct {
	Message string `json:"message"`
}

func ErrorJsonStructResponse(message string) *ErrorWithJSONFormat {
	return &ErrorWithJSONFormat{
		Message: message,
	}
}
