package main

import (
	"bytes"
	"testing"
)

func TestHello(t *testing.T) {
	// create buffer of bytes to simulate io.Writer
	buffer := bytes.Buffer{}

	// the msg to be passed to the function
	msg := "Hello World"
	Hello(&buffer, msg)

	got := buffer.String()
	want := msg

	if got != want {
		t.Errorf("got '%s' want '%s'")
	}
}
