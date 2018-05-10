package main

import "fmt"

func main() {

	var i uint8
	i = 20

	var f float32
	f = 214.437

	var b bool
	b = true

	var msg string
	msg = "hello"

	type myStringT string
	var myMessage myStringT
	myMessage = "special message"

	fmt.Printf("%d %v %T \n", i, i, i)
	fmt.Printf("%5.1f %.4f %g %T \n", f, f, f, f)
	fmt.Printf("%t %v %T \n", b, b, b)
	fmt.Printf("%s %T \n", msg, msg)
	fmt.Printf("%v %T \n", myMessage, myMessage)
}
