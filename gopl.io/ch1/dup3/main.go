package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

// Dup function takes a slice of byte and prints out strings
// which occur more than once in it.
func Dup(data []byte, out io.Writer) {
	counts := make(map[string]int)

	for _, line := range strings.Split(string(data), "\n") {
		counts[line]++
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Fprintf(out, "%s\t%d\n", line, n)
		}
	}
}

func main() {
	for _, filename := range os.Args[1:] {
		data, err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup: %v\n", err)
			continue
		}
		Dup(data, os.Stdout)
	}
}
