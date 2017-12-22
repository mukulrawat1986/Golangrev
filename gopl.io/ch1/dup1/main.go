// dup1 prints the text of each line that appears more than
// once in the standard input, preceeded by its count.
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// make a map of string to int
	counts := make(map[string]int)

	// create a scanner object on the standard input
	input := bufio.NewScanner(os.Stdin)

	// run a loop over the input
	for input.Scan() {
		counts[input.Text()]++
	}

	// Note: ignoring potential error from input.Err()
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
