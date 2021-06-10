package main

import (
	"fmt"

	"example.com/greetings"
)

func main() {
    // fmt.Println("Hello, World!")
	// fmt.Println(quote.Go())
	message := greetings.Hello("Haasin")
	fmt.Println(message)
}