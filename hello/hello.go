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

	// request a greeting msg
	message, err := greetings.Hello("")
	// if error print to console and exit program
	if (err != nil) {
		log.Fatal(err)
	} 

	fmt.Println(message)
}