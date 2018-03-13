package main

import (
	"bytes"
	"testing"
)

func TestEcho(t *testing.T) {
	t.Run("when arguments are passed", func(t *testing.T) {
		// create a buffer of bytes.Buffer which we will inject as our
		// dependency for printing.
		buffer := bytes.Buffer{}

		// a slice to mock our command line arguments
		args := []string{"filename", "Hello", "World"}

		Echo(&buffer, args)

		got := buffer.String()
		want := args[1] + " " + args[2] + "\n"
		if got != want {
			t.Errorf("got %#v want %#v", got, want)
		}
	})
}
