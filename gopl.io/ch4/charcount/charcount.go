// Charcount computes the counts of Unicode characters
package main

import "unicode/utf8"

func main() {
	counts := make(map[rune]int)    // counts of Unicode characters
	var utflen [utf8.UTFMax + 1]int // count of length of UTF-8 encoding
	invalid := 0                    // count of invalid UTF-8 characters
}
