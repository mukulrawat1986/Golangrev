package main

import (
	"fmt"
	"io"
	"strings"
)

// Echo prints the command line arguments
func Echo(w io.Writer, args []string) {
	fmt.Fprintln(w, strings.Join(args[1:], " "))
}

func main() {

}
