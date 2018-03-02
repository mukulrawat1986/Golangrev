package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func Echo(w io.Writer, args []string) {
	res := strings.Join(args[1:], " ")
	fmt.Fprintln(w, res)
}

func main() {
	Echo(os.Stdout, os.Args)
}
