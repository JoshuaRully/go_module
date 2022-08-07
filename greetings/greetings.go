package greetings

import (
	"fmt"
	"errors"
	"math/rand"
	"time"
)

// Hello returns a greeting for the named person
func Hello(name string) (string, error) {
	// if no name given throw error
	if name == "" {
		return "", errors.New("empty name")
	}

	// return a greeting that embeds the name in a msg
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

// Hellos returns a map that associates each of the named people
// with a greeting msg
func Hellos (names []string) (map[string]string, error) {
	// a map to associate names with msgs
	messages := make(map[string]string)
	// loop through received slice of names, calling Hello func to get
	// a msg for each name
	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}
		// in the map, associated the retrieved msf with the name
		messages[name] = message
	}
	return messages, nil
}

// init sets initial values for vars used in the func
func init() {
	rand.Seed(time.Now().UnixNano())
}

// randomFormat returns one of a set of random greeting msgs
func randomFormat() string {
	// a slice of msg formats
	formats := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v!",
		"Hail, %v! Well met!",
	}

	// return a randomly selected msf format by specifying random idx
	return formats[rand.Intn(len(formats))]
}