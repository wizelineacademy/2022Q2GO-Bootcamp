package errors

// ServiceError should be used to return error messages in JSON format
type ErrorMessage struct {
	Message string `json:"message"`
}
