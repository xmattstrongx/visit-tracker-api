package services

import "fmt"

type CityNotExistInStateError struct {
	message string
}

func newCityNotExistInStateError(city, state string) error {
	return CityNotExistInStateError{message: fmt.Sprintf("%s does not exist in state: %s", city, state)}
}

func (c CityNotExistInStateError) Error() string {
	return c.message
}
