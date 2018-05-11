package main

import "fmt"

func main() {

	s := "界"

	bytes := []byte(s)

	fmt.Println(bytes)
	fmt.Printf("%U \n", bytes)

	for _, b := range bytes {
		fmt.Printf("%d ", b)
	}
	fmt.Println()

	for _, b := range bytes {
		fmt.Printf("%o ", b)
	}
	fmt.Println()

	for _, b := range bytes {
		fmt.Printf("%x ", b)
	}
	fmt.Println()

	fmt.Println()
	fmt.Println("\xf0\xa0\x9c\x8e")
	fmt.Println("\360\240\234\216")

	ch := "界"
	fmt.Printf("%q \n", ch)
	fmt.Print("%+q \n", ch)

	fmt.Printf("\U0002070e")

	fmt.Println(string(0x0002070e))
}
