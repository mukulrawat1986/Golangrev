package main

import (
	"fmt"
	"io"
	"os"
)

// Hello function greets the name passed.
func Hello(buffer io.Writer, name string) {
	if name == "" {
		fmt.Fprintln(buffer, "Hello World!")
	} else {
		fmt.Fprintf(buffer, "Hello %s\n", name)
	}
}

func main() {
	Hello(os.Stdout, "World")
}
