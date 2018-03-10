package main

import (
	"fmt"
)

// You need to add channel to cancel this request
func Poke(team Team) {
	// Send a request to the helthcheck endpoint
	// Add history if it needs
	// Update States if necessary
	// Emit a log
	fmt.Println("I'm poking")
}
