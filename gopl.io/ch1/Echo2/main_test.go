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
		want := args[0] + " " + args[1] + "\n"

		if got != want {
			t.Errorf("got '%s' want '%s'", got, want)
		}
	})
}
