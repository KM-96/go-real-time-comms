package reader

import (
	"bufio"
	"fmt"
	"os"
)

type Reader interface {
	Read() string
}

type ConsoleReader struct{}

// Create a new scanner to read input from the console
var scanner = bufio.NewScanner(os.Stdin)

func (cr *ConsoleReader) Read() string {
	fmt.Print("Enter your message: ")

	// Read user input
	scanner.Scan()
	input := scanner.Text()

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading input:", err)
		panic("Error reading input from console")
	}
	return input
}
