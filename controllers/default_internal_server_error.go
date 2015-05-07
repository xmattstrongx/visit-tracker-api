package controllers

import (
	"fmt"
	"log"
)

type DefaultInternalServerError struct {
	message string
}

func newDefaultInternalServerError(err error) error {
	log.Printf("Unhandled general error: %v", err)
	return DefaultInternalServerError{message: fmt.Sprintf("Sorry we are. Occured an error seems to have...")}
}

func (d DefaultInternalServerError) Error() string {
	return d.message
}
