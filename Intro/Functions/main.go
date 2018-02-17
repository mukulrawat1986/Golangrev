package main

import "os"

func printer(msg string) error {
	f, err := os.Create("helloworld.txt")
	if err != nil {
		return err
	}
	defer f.Close()
	f.Write([]byte(msg))
	return nil
}

func main() {
	printer("Hello, World!")
}
