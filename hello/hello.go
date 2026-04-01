package main

import (
	"fmt"

	"github.com/gokuljs/greetings"
)

func main() {
	var message string
	message = greetings.Hello("Gokul js")
	fmt.Println(message)
}
