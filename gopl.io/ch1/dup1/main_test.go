package main

import (
	"bytes"
	"testing"
)

func TestDup(t *testing.T) {
	// create a buffer of bytes.Buffer where the output is stored
	outBuffer := &bytes.Buffer{}

	// create the input string
	input := "Hello\nhello\nHello World\nhello\nHello\nhello World\n"

	// create an input buffer and initialize it with the input string
	inputBuffer := bytes.NewBufferString(input)

	Dup(inputBuffer, outBuffer)

	got := outBuffer.String()
	want := "Hello\t2\nhello\t2\n"

	if got != want {
		t.Errorf("got '%#v' want '%#v'", got, want)
	}
}
