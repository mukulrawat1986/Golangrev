package main

import (
	"bytes"
	"fmt"
)

// Greet function writes a greeting to the writer
func Greet(writer *bytes.Buffer, str string) {
	fmt.Fprintf(writer, "%s\n", str)
}

func main() {

}
