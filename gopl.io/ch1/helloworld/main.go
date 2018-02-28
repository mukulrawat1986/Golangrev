package main

import (
	"fmt"
	"io"
	"os"
)

func Hello(w io.Writer, msg string) {
	// fmt.Fprintf(w, "%s\n", msg)
	fmt.Fprintln(w, msg)
}

func main() {
	Hello(os.Stdout, "Hello, World!")
}
