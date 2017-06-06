package main

import (
	"errors"
	"fmt"
	"os"
)

var (
	errorEmptyString = errors.New("Unwilling to print an empty string")
)

func printer(msg string) error {
	if msg == "" {
		return errorEmptyString
	}
	_, err := fmt.Printf("%s\n", msg)
	return err
}

func main() {
	if err := printer("Hello, World!!"); err != nil {
		if err == errorEmptyString {
			fmt.Printf("You tried to print an empty string!")
		} else {
			fmt.Printf("Printer failed: %s\n", err)
		}
		os.Exit(1)
	}
}
