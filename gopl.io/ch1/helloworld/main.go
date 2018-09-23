package main

import (
	"fmt"
	"io"
	"os"
)

func Hello(w io.Writer, name string) {
	if len(name) == 0 {
		name = "World"
	}
	fmt.Fprintf(w, "Hello, %s!\n", name)
}

func main() {
	Hello(os.Stdout, "")
}
