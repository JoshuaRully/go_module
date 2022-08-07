package main

import (
	"fmt"
	"log"
	"example.com/greetings"
)

func main() {
	// set props of predefined logger including log entry prefix
	// and a flag to disable printing the time, source file and line number
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	// a slice of names
	names := []string{"Josh", "Jessica", "Jebediah"}

	// request greeting msgs for the names
	messages, err := greetings.Hellos(names)
	// if error print to console and exit program
	if (err != nil) {
		log.Fatal(err)
	}

	// print returned map of msgs to console
	fmt.Println(messages)
}