package errorext

import "fmt"

type ServiceError struct {
	Message string
}

func (e ServiceError) Error() string {
	return e.Message
}

func NewServiceError(format string, a ...any) *ServiceError {
	return &ServiceError{
		Message: fmt.Sprintf(format, a...),
	}
}
