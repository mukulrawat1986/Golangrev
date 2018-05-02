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
}
