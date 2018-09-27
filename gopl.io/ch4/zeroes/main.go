// zeroes the content of a 32 byte array
package main

import "fmt"

func main() {
	var x [32]byte
	for i := 0; i < 32; i++ {
		x[i] = 'C'
	}
	fmt.Printf("%s\n", x)
	zero(&x)
	fmt.Printf("%v\n", x)
}

func zero(ptr *[32]byte) {
	for i := range ptr {
		ptr[i] = 0
	}
}
