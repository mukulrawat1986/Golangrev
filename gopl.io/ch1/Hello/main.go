// Program to print Hello World
package main

import (
	"fmt"
	"io"
	"os"
)

func Hello(w io.Writer, s string) {
	fmt.Fprintf(w, "%s\n", s)
}

func main() {
	Hello(os.Stdout, "Hello World")
}
