package main

import (
	"fmt"

	"codebyte.cy/greetings"
)

func main() {
	// Call Hello and handle its return values
	message, err := greetings.Hello("Stefanos")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Message:", message)
}
