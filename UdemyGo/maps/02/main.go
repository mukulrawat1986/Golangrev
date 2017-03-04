package main

import "fmt"

func main() {

	myGreeting := map[int]string{
		0: "Good Morning!",
		1: "Bonjour",
		2: "Bueno dias",
		3: "Bongiorno",
	}

	fmt.Println(myGreeting)

	if val, ok := myGreeting[2]; ok {
		delete(myGreeting, 2)
		fmt.Println("val: ", val)
		fmt.Println("ok: ", ok)
	} else {
		fmt.Println("That value doesn't exists.")
		fmt.Println("val: ", val)
		fmt.Println("ok: ", ok)
	}

	fmt.Println(myGreeting)
}
