// basename removes directory component and a . suffix.
// eg. a => a, a.go => a, a/b/c.go => c, a/b.c.go => b.c
package main

import "fmt"

func basename(s string) string {
	// Discard last / and everything before it
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '/' {
			s = s[i+1:]
			break
		}
	}

	// preserve everything before last '.'
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '.' {
			s = s[:i]
			break
		}
	}
	return s
}


func main() {
	fmt.Println(basename("a"))
	fmt.Println(basename("a/b/c.go"))
	fmt.Println(basename("a.go"))
	fmt.Println(basename("a/b.c.go"))
}

