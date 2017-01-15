// program to zero out content of an array

package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	x := sha256.Sum256([]byte("x"))
	fmt.Printf("%x\n", x)
	zero1(&x)
	fmt.Printf("%x\n", x)
	x = sha256.Sum256([]byte("y"))
	fmt.Printf("%x\n", x)
	zero2(&x)
	fmt.Printf("%x\n", x)
}

func zero1(ptr *[32]byte) {
	for i := range ptr {
		ptr[i] = 0
	}
}

func zero2(ptr *[32]byte) {
	*ptr = [32]byte{}
}
