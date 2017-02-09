package main

import "fmt"

func printer(msg string) (string, error) {
	msg += "\n"
	_, err := fmt.Printf("%s", msg)
	return msg, err
}

func main() {
	appendedMessage, err := printer("Hello, World!")
	if err == nil {
		fmt.Printf("%q\n", appendedMessage)
	}
}
