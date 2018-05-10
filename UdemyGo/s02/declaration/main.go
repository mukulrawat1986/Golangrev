package main

import "fmt"

func main() {

	s1, s2 := "msg1", "msg2"
	var s3 = "msg3"
	var s4 string = "msg4" // not the best way to do it, should be avoided

	fmt.Println(s1 + " " + s2 + " " + s3 + " " + s4)
}
