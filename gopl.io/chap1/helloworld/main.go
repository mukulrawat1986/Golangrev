package main

import (
	"bytes"
	"fmt"
)

// Hello function greets a named string
func Hello(writer *bytes.Buffer, name string) {
	fmt.Fprintf(writer, "Hello, %s!\n", name)
}

func main() {

}
