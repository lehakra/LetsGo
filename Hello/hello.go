package main

import (
	"fmt"
	"log"

	"romil.co/greetings"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	names := []string{"Romil", "Ghanshyam"}

	messages, err := greetings.Hellos(names)

	if err != nil {
		log.Fatal(err) // this is exiting the program? wow
	}

	fmt.Println(messages)
}
