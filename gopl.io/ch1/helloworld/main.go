package main

import (
	"bytes"
	"fmt"
)

// Hello function greets the person or prints out "Hello World"
func Hello(b *bytes.Buffer, name string) {
	fmt.Fprintf(b, "Hello %s!\n", name)
}

func main() {

}
