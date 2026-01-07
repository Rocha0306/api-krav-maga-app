package Presentation

import (
	"errors"
)

func ThrowException(error_message string) error {
	return errors.New(error_message)
}
