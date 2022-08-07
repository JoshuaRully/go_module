package greetings

import (
	"fmt"
	"errors"
)

// Hello returns a greeting for the named person
func Hello(name string) (string, error) {
	// if no name given throw error
	if name == "" {
		return "", errors.New("empty name")
	}

	// return a greeting that embeds the name in a msg
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message, nil
}