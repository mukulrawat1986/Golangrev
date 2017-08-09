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
	// close the file when done
	defer f.Close()

	// Read from the file using a byte slice
	b := make([]byte, 100)

	n, err := f.Read(b)

	fmt.Printf("%d: % x\n", n, b)

	stringVersion := string(b)
	fmt.Printf("%d: %s\n", n, stringVersion)

	// converting a string to file

	someString := "foo bar"

	n, _ = f.Write([]byte(someString))
	fmt.Printf("%d\n", n)
}
