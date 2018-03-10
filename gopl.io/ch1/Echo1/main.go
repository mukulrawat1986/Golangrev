// Echo1 prints its command line arguments
package main

import (
	"fmt"
	"io"
	"os"
)

func Echo(w io.Writer, args []string) {
	var s, sep string
	for i := 1; i < len(args); i++ {
		s += sep + args[i]
		sep = " "
	}
	fmt.Fprintf(w, "%s\n", s)
}

func main() {
	Echo(os.Stdout, os.Args)
}