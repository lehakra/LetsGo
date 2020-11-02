package main

import (
	"fmt"

	"romil.co/greetings"
)

func main() {
	message := greetings.Hello("Romil")
	fmt.Println(message)
}
