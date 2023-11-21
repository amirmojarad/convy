package errorext

import "fmt"

type DatabaseError struct {
	Message string
}

func (e DatabaseError) Error() string {
	return e.Message
}

func NewDatabaseError(format string, a ...any) *DatabaseError {
	return &DatabaseError{
		Message: fmt.Sprintf(format, a...),
	}
}
