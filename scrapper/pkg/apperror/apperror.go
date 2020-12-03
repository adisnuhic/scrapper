package apperror

// AppError cusom defined app error
type AppError struct {
	Code    uint64
	Cause   error
	Message string
}

// New returns new error object
func New(code uint64, err error, message string) *AppError {
	return &AppError{
		Code:    code,
		Cause:   err,
		Message: message,
	}
}
