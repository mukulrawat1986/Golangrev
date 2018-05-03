package main

import (
	"bytes"
	"testing"
)

func TestEcho(t *testing.T) {

	t.Run("when arguments are passed", func(t *testing.T) {
		// create a buffer
		buffer := bytes.Buffer{}

		// arguments
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

	t.Run("when no arguments are passed", func(t *testing.T) {
		// create the buffer
		buffer := bytes.Buffer{}

		// create the args
		args := []string{
			"filename",
		}

		Echo(&buffer, args)

		got := buffer.String()
		want := "\n"

		if got != want {
			t.Errorf("got '%s' want '%s'", got, want)
		}
	})
}
