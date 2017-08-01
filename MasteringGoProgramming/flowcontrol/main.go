package main

import (
	"fmt"
	"runtime"
)

func main() {

	// defer
	defer fmt.Println("Exiting main function.......")
	fmt.Println("Entering main function")

	// if statement
	inc := increment()

	if i := inc(); i < 0 {
		fmt.Println("i is a negative number")
	} else if i == 0 {
		fmt.Println("i is zero")
	} else {
		fmt.Println("i is a positive number")
	}

	// switch statement
	switch i := inc(); {
	case i < 0:
		fmt.Println("i is a negative number")
	case i == 0:
		fmt.Println("i is zero")
	default:
		fmt.Println("i is a positive number")
	}

	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows..
		fmt.Printf("%s.", os)
	}
}

func increment() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}
