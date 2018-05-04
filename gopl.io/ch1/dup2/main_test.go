package main

import (
	"bytes"
	"testing"
)

func TestDup(t *testing.T) {
	// create a buffer to store output of type bytes.Buffer
	outBuffer := &bytes.Buffer{}

	// create an input string
	input := "Hello\nhello\nHello\nhello world\nHello\nhello\n"

	// create an input buffer using the input string
	inBuffer := bytes.NewBufferString(input)

	Dup(inBuffer, outBuffer)

	got := outBuffer.String()
	want := "Hello\t3\nhello\t2\n"

	if got != want {
		t.Errorf("got '%#v' want '%#v'", got, want)
	}
}
