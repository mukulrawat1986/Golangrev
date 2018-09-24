package main

import (
	"fmt"
	"io"
)

// Echo function prints its command line arguments
func Echo(w io.Writer, args []string) {
	s, sep := "", ""
	for _, arg := range args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Fprintln(w, s)
}

func main() {
}
