package main

import (
	"bytes"
	"testing"
)

func TestHello(t *testing.T) {

	t.Run("when a name is sent", func(t *testing.T) {
		// set up a buffer of bytes.Buffer
		// Test function will print here
		buffer := bytes.Buffer{}

		Hello(&buffer, "Chris")

		got := buffer.String()
		want := "Hello Chris\n"

		if got != want {
			t.Errorf("got '%#v' want '%#v'", got, want)
		}
	})

	t.Run("when no name is passed", func(t *testing.T) {
		// create a buffer of bytes.Buffer
		buffer := bytes.Buffer{}

		Hello(&buffer, "")

		got := buffer.String()
		want := "Hello World!\n"

		if got != want {
			t.Errorf("got '%#v' want '%#v'", got, want)
		}
	})
}
