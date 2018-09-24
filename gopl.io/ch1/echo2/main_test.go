package main

import (
	"bytes"
	"testing"
)

func TestEcho(t *testing.T) {

	// There are two dependecies in the Echo function. First the command line arguments
	// and secondly the dependency of printing. We will pass a slice of strings to mock the
	// command line arguments, and pass an interface to record what is being printed.
	args := []string{"filename", "Hello", "World"}

	buffer := &bytes.Buffer{}

	Echo(buffer, args)

	got := buffer.String()
	want := "Hello World\n"

	if got != want {
		t.Errorf("got '%#v', want '%#v'", got, want)
	}
}
