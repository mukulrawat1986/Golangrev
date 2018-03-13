// Echo prints its command line arguments.
package main

import (
	"fmt"
	"io"
)

// Echo prints the command line arguments passed to it
func Echo(writer io.Writer, args []string) {
	var s, sep string
	for i := 1; i < len(args); i++ {
		s += sep + args[i]
		sep = " "
	}
	fmt.Fprintf(writer, "%s\n", s)
}

func main() {

}
