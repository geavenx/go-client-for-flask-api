package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please provide a command as an argument.")
		return
	}

	command := os.Args[1]

	switch command {
	case "sum":
		Sum()
	case "subtract":
		Subtract()
	case "multiply":
		Multiply()
	case "divide":
		Divide()
	case "hello":
		fmt.Println("hello")
	}
}
