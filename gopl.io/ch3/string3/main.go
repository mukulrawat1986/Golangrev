package main

import "fmt"

func main() {
	// hasPrefix
	fmt.Println(hasPrefix("abcdef", "abc"))
	fmt.Println(hasPrefix("abcdef", "def"))

	// hasSUffix
	fmt.Println(hasSuffix("armda", "da"))
	fmt.Println(hasSuffix("asrm", "asrm"))
	fmt.Println(hasSuffix("armds", "md"))

	// containsSubString
	fmt.Println(containsSubString("hello", "el"))
	fmt.Println(containsSubString("hello world this is dog", "his"))
}

func hasPrefix(s, prefix string) bool {
	return len(s) >= len(prefix) && s[:len(prefix)] == prefix
}

func hasSuffix(s, suffix string) bool {
	return len(s) >= len(suffix) && s[len(s)-len(suffix):] == suffix
}

func containsSubString(s, substr string) bool {
	for i := 0; i < len(s); i++ {
		if hasPrefix(s[i:], substr) {
			return true
		}
	}
	return false
}
