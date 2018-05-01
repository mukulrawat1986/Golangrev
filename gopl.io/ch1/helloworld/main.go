package main

import (
	"fmt"
	"io"
	"os"
)

// Hello function greets the person or prints out "Hello World"
func Hello(w io.Writer, name string) {
	fmt.Fprintf(w, "Hello %s!\n", name)
}

func main() {
	Hello(os.Stdout, "")
}
