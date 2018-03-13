package main

import (
	"fmt"
	"io"
	"os"
)

var (
	greeting = "Hello, World!"
)

// Greet function writes a greeting to the writer
func Greet(writer io.Writer, str string) {
	fmt.Fprintf(writer, "%s\n", str)
}

func main() {
	Greet(os.Stdout, greeting)
}
