package main

import "fmt"

func main() {

	i := 38

	fmt.Printf("%7b %#3o %d %#x %s \n", i, i, i, i, string(i))

	i = 40
	fmt.Printf("%7b %#3o %d %#x %s \n", i, i, i, i, string(i))

	ch := 'e'
	fmt.Printf("\n%T\n", ch)
	fmt.Printf("%7b %#3o %d %#x %s \n", ch, ch, ch, ch, string(ch))

	ch = 'E'
	fmt.Printf("\n%T\n", ch)
	fmt.Printf("%7b %#3o %d %#x %s \n", ch, ch, ch, ch, string(ch))

	i = 958
	fmt.Printf("%7b %#3o %d %#x %s \n", i, i, i, i, string(i))

	ch = 'P'
	fmt.Printf("\n%T\n", ch)
	fmt.Printf("%7b %#3o %d %#x %s \n", ch, ch, ch, ch, string(ch))

	ch = 'A'
	fmt.Printf("%v \t %v \t\t\t %b\n", string(ch), []byte(string(ch)), ch)

	ch = '$'
	fmt.Printf("%v \t %v \t\t\t %b\n", string(ch), []byte(string(ch)), ch)

	ch = 'こ'
	fmt.Printf("%v \t %v \t\t\t %b\n", string(ch), []byte(string(ch)), ch)

	ch = '世'
	fmt.Printf("%v \t %v \t\t\t %b\n", string(ch), []byte(string(ch)), ch)
}
