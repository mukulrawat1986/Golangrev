package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// Dup prints the count and text of lines that appear more than once
// in the input. It reads from stdin or a list of named files.
func Dup(i io.Reader, w io.Writer) {
	counts := make(map[string]int)
	input := bufio.NewScanner(i)
	for input.Scan() {
		counts[input.Text()]++
	}
	if input.Err() != nil {
		fmt.Fprintf(os.Stderr, "dup: %v\n", input.Err())
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Fprintf(w, "%s\t%d\n", line, n)
		}
	}
}

func main() {

}
