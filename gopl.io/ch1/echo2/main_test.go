package main

import (
	"bytes"
	"testing"
)

func TestEcho(t *testing.T) {
	// create a buffer of bytes.Buffer which we will use to inject
	// as our dependency of printing.
	buffer := bytes.Buffer{}

	// mock our command line arguments as a slice
	args := []string{"filename", "Hello", "World"}

	Echo(&buffer, args)

	got := buffer.String()
	want := args[1] + " " + args[2] + "\n"
	if got != want {
		t.Errorf("got %#v want %#v", got, want)
	}
}
