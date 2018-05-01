package main

import (
	"bytes"
	"testing"
)

func TestHello(t *testing.T) {
	// create a buffer using bytes.Buffer
	buffer := bytes.Buffer{}

	Hello(&buffer, "Mukul")

	got := buffer.String()
	want := "Hello Mukul!\n"

	if got != want {
		t.Errorf("want '%#v' got '%#v'", want, got)
	}
}
