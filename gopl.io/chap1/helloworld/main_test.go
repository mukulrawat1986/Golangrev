package main

import (
	"bytes"
	"testing"
)

func TestHello(t *testing.T) {
	// create a buffer of bytes.Buffer to store the output.
	// this buffer is passed as an argument in the test call.
	buffer := bytes.Buffer{}

	// Hello function called
	Hello(&buffer, "Chris")

	got := buffer.String()
	want := "Hello, Chris!\n"

	if got != want {
		t.Errorf("got %#v want %#v", got, want)
	}
}
