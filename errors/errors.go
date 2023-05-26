package errors

type AppError struct {
	Message map[string]string `json:"message"`
}

func NewAppError(message map[string]string) *AppError {

	return &AppError{
		Message: message,
	}
}
