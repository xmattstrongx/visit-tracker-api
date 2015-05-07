package repositories

import (
	"fmt"
	"log"
)

type ConnectionTimeoutError struct {
	message string
}

func newConnectionTimeoutError() error {
	log.Print("Timed out connecting to the database")
	return ConnectionTimeoutError{message: fmt.Sprintf("Timed out connecting to the database")}
}

func (c ConnectionTimeoutError) Error() string {
	return c.message
}
