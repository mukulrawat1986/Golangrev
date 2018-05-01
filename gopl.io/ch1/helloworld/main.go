package main

import (
	"fmt"
	"io"
	"os"
)

// Hello function greets the person or prints out "Hello World"
func Hello(w io.Writer, name string) {
	if len(name) != 0 {
		fmt.Fprintf(w, "Hello %s!\n", name)
	} else {
		fmt.Fprintln(w, "Hello, World!")
	}
}

func main() {
	Hello(os.Stdout, "")
}
