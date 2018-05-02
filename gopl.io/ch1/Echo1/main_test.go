package main

import (
	"bytes"
	"testing"
)

func TestEcho(t *testing.T) {

	t.Run("when command line arguments are passed", func(t *testing.T) {
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
	})

	t.Run("when command line arguments are not passed", func(t *testing.T) {
		// create a buffer of bytes.Buffer
		buffer := bytes.Buffer{}

		args := []string{"filename"}

		Echo(&buffer, args)

		got := buffer.String()
		want := "\n"

		if got != want {
			t.Errorf("got '%s' want '%s'", got, want)
		}
	})
}
