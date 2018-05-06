package main

import (
	"fmt"
	"io"
	"os"
)

// Hello function greets a named string
func Hello(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello, %s!\n", name)
}

func main() {
	Hello(os.Stdout, "Chris")
}
