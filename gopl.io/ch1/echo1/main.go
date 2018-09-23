package main

import (
	"fmt"
	"io"
	"os"
)

// Echo prints its command line arguments
func Echo(w io.Writer, args []string) {
	var s, sep string
	for i := 1; i < len(args); i++ {
		s += sep + args[i]
		sep = " "
	}
	fmt.Fprintln(w, s)
}

func main() {
	Echo(os.Stdout, os.Args)
}
