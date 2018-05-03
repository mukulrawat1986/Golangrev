package main

import (
	"bytes"
	"testing"
)

func TestEcho(t *testing.T) {

	assertResult := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got '%#v' want '%#v'", got, want)
		}
	}

	t.Run("when arguments are given", func(t *testing.T) {
		// create a buffer
		buffer := bytes.Buffer{}

		// create the arguments
		args := []string{
			"filename",
			"Hello",
			"World",
		}

		Echo(&buffer, args)

		got := buffer.String()
		want := args[1] + " " + args[2] + "\n"

		assertResult(t, got, want)
	})

	t.Run("when no arguments are given", func(t *testing.T) {
		// create a buffer
		buffer := bytes.Buffer{}

		// create a slice of arguments
		args := []string{"filename"}

		Echo(&buffer, args)

		got := buffer.String()
		want := "\n"

		assertResult(t, got, want)
	})
}
