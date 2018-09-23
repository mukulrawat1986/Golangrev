package main

import (
	"bytes"
	"testing"
)

func TestHello(t *testing.T) {

	t.Run("when we pass a name to the Hello function", func(t *testing.T) {
		// There is a dependency of printing in Hello function. So to capture that, we pass
		// an interface, thus we can capture the output that is printed there. In our main function
		// we pass the stdout
		buffer := &bytes.Buffer{}

		Hello(buffer, "Chris")

		got := buffer.String()
		want := "Hello, Chris!\n"

		if got != want {
			t.Errorf("wanted '%#v', got '%#v'", want, got)
		}
	})
}
