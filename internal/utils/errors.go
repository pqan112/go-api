package utils

const (
	ErrCodeBadRequest = "BAD_REQUEST"
	ErrCodeNotFound   = "NOT_FOUND"
	ErrCodeConflict   = "CONFLICT"
	ErrCodeInternal   = "INTERNAL"
)

type AppError struct {
	Message    string `json:"message"`
	StatusCode int    `json:"statusCode"`
	Err        error  `json:"error"`
}

type ErrorDetail struct {
	Path    string `json:"path"`
	Message string `json:"message"`
}

type ValidationErrorResponse struct {
	Message    []ErrorDetail `json:"message"`
	Error      string        `json:"error"`
	StatusCode int           `json:"statusCode"`
}

func (ae *AppError) Error() string {
	return ""
}

func NewError(message string, code int) error {
	return &AppError{
		Message:    message,
		StatusCode: code,
	}
}

func WrapError(message string, code int, err error) error {
	return &AppError{
		Message:    message,
		StatusCode: code,
		Err:        err,
	}
}
