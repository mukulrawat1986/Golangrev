package main

import (
	"bytes"
	"fmt"
)

// Hello function greets the name passed.
func Hello(buffer *bytes.Buffer, name string) {
	fmt.Fprintf(buffer, "Hello %s\n", name)
}

func main() {

}
