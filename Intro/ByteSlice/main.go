package main

import (
	"fmt"
	"os"
)

func main() {
	// open a file
	f, err := os.Open("test.txt")
	if err != nil {
		fmt.Printf("%s\n", err)
		os.Exit(1)
	}
	defer f.Close()

	// byte slice of 100 length
	b := make([]byte, 100)

	n, err := f.Read(b)
	if err != nil {
		fmt.Printf("%s\n", err)
		os.Exit(1)
	}
	fmt.Printf("%d: %s\n", n, string(b))
}
