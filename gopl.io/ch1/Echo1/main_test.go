package main

import (
	"bytes"
	"testing"
)

func TestEcho(t *testing.T) {
	// create a buffer using bytes.Buffer
	buffer := bytes.Buffer{}

	// use a slice to store the arguments
	args := []string{
		"filename",
		"Hello",
		"World",
	}

	Echo(&buffer, args)

	got := buffer.String()
	want := args[1] + " " + args[2] + "\n"

	if got != want {
		t.Errorf("got '%#v' want '%#v'", got, want)
	}
}
