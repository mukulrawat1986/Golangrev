package main

import "fmt"

func printer(msg string) (e error) {
	_, e = fmt.Printf("%s\n", msg)
	return
}

func main() {
	printer("Hello, World!")
}
