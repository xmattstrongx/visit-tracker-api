package services

import "fmt"

type VisitValidationError struct {
	message string
}

func newVisitValidationError(value string) error {
	return VisitValidationError{message: fmt.Sprintf("Request missing value: %s", value)}
}

func (c VisitValidationError) Error() string {
	return c.message
}
