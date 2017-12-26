// Dup3 works in non-streaming mode, reads all the input into memory
// and then print out count and text of lines that appear more than once.
package main

import (
	"fmt"
	"strings"
	"io/ioutil"
	"os"
)

func main()  {
	counts := make(map[string]int)
	for _, filename := range os.Args[1:] {
		data, err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup3: %v\n", err)
			continue
		}
		for _, line := range strings.Split(string(data), "\n") {
			counts[line]++
		}

		for line, n := range counts {
			if n > 1 {
				fmt.Printf("%d\t%s\n", n, line)
			}
		}
	}
}