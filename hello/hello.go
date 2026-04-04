package main

import (
	"fmt"
	"log"

	"github.com/gokuljs/greetings"
)

func main() {
	log.SetPrefix("greetings: ")
	names := []string{"gokul", "matt", "michael"}
	messages, err := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(messages)
}
