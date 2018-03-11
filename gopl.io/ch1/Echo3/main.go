// Echo3 prints its command-line arguments
package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func Echo(w io.Writer, args []string) {
	fmt.Fprintf(w, "%s\n", strings.Join(args[1:], " "))
}

func main() {
	Echo(os.Stdout, os.Args)
}
