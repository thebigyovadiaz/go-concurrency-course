package main

import (
	"fmt"
	"time"
)

func main() {
	go hello()

	// It's not recommend use sleep in production
	time.Sleep(1 * time.Second)

	goodbye()
}

func hello() {
	fmt.Println("Hello, world!")
}

func goodbye() {
	fmt.Println("Goodbye, world!")
}
