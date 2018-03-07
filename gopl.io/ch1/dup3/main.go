package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	counts := make(map[string]int)
	for _, filename := range os.Args[1:] {
		data, err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup3: %v\n", err)
			continue
		}
		Count(data, counts)
	}

	Print(os.Stdout, counts)
}

func Count(data []byte, counts map[string]int) {
	for _, line := range strings.Split(string(data), "\n") {
		if line != "" {
			counts[line]++
		}
	}
}

func Print(w io.Writer, counts map[string]int) {
	for line, n := range counts {
		if n > 1 {
			fmt.Fprintf(w, "%d\t%s\n", n, line)
		}
	}
}
