package main

import "fmt"

func printer(format string, msgs ...string) {
	for _, msg := range msgs {
		fmt.Printf(format, msg)
	}
}

func main() {
	printer("%s\n", "Hello", "world", "this", "is", "dog")
}
