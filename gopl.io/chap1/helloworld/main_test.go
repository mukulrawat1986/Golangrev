package main

import (
	"bytes"
	"testing"
)

func TestHello(t *testing.T) {

	t.Run("when a name is given", func(t *testing.T) {
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
	})

	t.Run("when name is an empty string", func(t *testing.T) {
		// create a buffer of bytes.Buffer
		buffer := bytes.Buffer{}

		Hello(&buffer, "")

		got := buffer.String()
		want := "Hello, World\n"

		if got != want {
			t.Errorf("got %#v want %#v", got, want)
		}
	})
}
