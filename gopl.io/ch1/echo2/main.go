package main

import (
	"fmt"
	"io"
	"os"
)

func Echo(w io.Writer, args []string) {
	s, sep := "", ""
	for _, arg := range args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Fprintln(w, s)
}

func main() {
	Echo(os.Stdout, os.Args)
}
