package main

import "fmt"

//  one string contains another as prefix
func hasPrefix(s, prefix string) bool {
	return len(s) >= len(prefix) && s[:len(prefix)] == prefix
}

// one string has another as suffix
func hasSuffix(s, suffix string) bool {
	return len(s) >= len(suffix) && s[len(s) - len(suffix):] == suffix
}

// one string contains another as substring
func contains(s, substr string) bool {
	for i := 0; i < len(s); i++ {
		if hasPrefix(s[i:], substr) {
			return true
		}
	}
	return false
}

func main() {

	// test prefix
	fmt.Println(hasPrefix("abcde", "bcde"))
	fmt.Println(hasPrefix("abc", "abc"))
	fmt.Println(hasPrefix("abcde", "abc"))
	fmt.Println(hasPrefix("abcde", "xabcde"))

}