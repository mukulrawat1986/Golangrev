package main

import "fmt"

func main() {
	var k1 uint8
	k1 = 20

	f1 := 3.14
	fmt.Printf("%f - %T\n", f1, f1)
	fmt.Printf("%5.3f %.2f \n", f1, 214.437)

	fmt.Println(k1)
}
