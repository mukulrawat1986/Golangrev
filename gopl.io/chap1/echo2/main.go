package main

import (
	"fmt"
	"io"
)

// Echo prints its command line arguments
func Echo(w io.Writer, args []string) {
	s, sep := "", ""
	for _, arg := range args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Fprintf(w, "%s\n", s)
}

func main() {

}
