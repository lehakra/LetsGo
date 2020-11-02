package greetings

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

// Hello returns greeting and error
func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("Empty Name")
	}

	message := fmt.Sprintf(randomFormat(), name) // := declared and initializes in one go.
	// message := fmt.Sprintf(randomFormat(), "") // := declared and initializes in one go.
	// alternative would have been to
	// var message string
	// message = fmt..
	return message, nil
}

// Hellos returns a map of greetings for each of the given names
func Hellos(names []string) (map[string]string, error) {
	messages := make(map[string]string) // slice map and chan only

	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}
		messages[name] = message
	}

	return messages, nil
}

func init() {
	rand.Seed(time.Now().UnixNano())
}
func randomFormat() string {
	formats := []string{
		"Hi, %v. Welcome!",
		"Great to see you %v",
		"Hail, %v! Well met!",
	}

	return formats[rand.Intn(len(formats))]
}
