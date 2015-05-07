package repositories

import (
	"fmt"
	"log"
)

type BadConnectionError struct {
	message string
}

func newBadConnectionError(err error) error {
	log.Printf("Bad connection error: %v", err)
	return BadConnectionError{message: fmt.Sprintf("Unable to connect to the DB")}
}

func (b BadConnectionError) Error() string {
	return b.message
}
