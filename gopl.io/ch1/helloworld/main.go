package main

import (
	"fmt"
	"io"
	"os"
)

// Hello functions greets the person whose name is passed as a string. If no name
// is provided the function prints out "Hello, World!"
func Hello(w io.Writer, name string) {
	if len(name) == 0 {
		name = "World"
	}
	fmt.Fprintf(w, "Hello, %s!\n", name)
}

func main() {
	Hello(os.Stdout, "")
}
