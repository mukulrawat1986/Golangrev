package main

import (
	"fmt"
	"os"
)

func printer(msg string) error {
	if msg == "" {
		return fmt.Errorf("Unwilling to print empty string")
	}
	_, err := fmt.Printf("%s\n", msg)
	return err
}

func main() {
	if err := printer("Hello, World!!"); err != nil {
		fmt.Printf("Printer failed: %s\n", err)
		os.Exit(1)
	}
}
