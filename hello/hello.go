package main

import (
	"fmt"
	"log"

	"github.com/gokuljs/greetings"
)

func main() {
	log.SetPrefix("greetings: ")
	var message string
	var err error
	message, err = greetings.Hello("")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(message)
}
