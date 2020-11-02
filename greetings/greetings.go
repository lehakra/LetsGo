package greetings

import "fmt"

// Hello returns greeting
func Hello(name string) string {
	message := fmt.Sprintf("Hi, %v. Welcome", name) // := declared and initializes in one go.
	// alternative would have been to
	// var message string
	// message = fmt..
	return message
}
