package main

import (
	"bytes"
	"testing"
)

func TestEcho(t *testing.T) {
	t.Run("when arguments are passed", func(t *testing.T) {
		// create a buffer to store the print output
		buffer := bytes.Buffer{}

		// use a slice of string to store the command line arguments
		args := []string{
			"filename",
			"Hello",
			"World",
		}

		// Echo function called
		Echo(&buffer, args)

		got := buffer.String()
		want := args[1] + " " + args[2] + "\n"

		if got != want {
			t.Errorf("got %#v want %#v", got, want)
		}
	})

	t.Run("when an empty string is passed as argument", func(t *testing.T) {
		// create a buffer to store the print output
		buffer := bytes.Buffer{}

		// use a slice of string to store the command line arguments
		args := []string{
			"filename",
		}

		// Echo function called
		Echo(&buffer, args)

		got := buffer.String()
		want := "\n"

		if got != want {
			t.Errorf("got %#v want %#v", got, want)
		}
	})
}
