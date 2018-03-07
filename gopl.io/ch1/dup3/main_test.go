package main

import (
	"bytes"
	"reflect"
	"testing"
)

func TestCount(t *testing.T) {
	// handle the dependency of input from a file
	// we can use a buffer of Bytes to represent it
	b := []byte("Hello\nHello\nHello World\nWorld\n")

	// create our buffer
	buffer := bytes.NewBuffer(b)

	// create our map
	counts := make(map[string]int)

	Count(buffer, counts)

	got := counts
	want := map[string]int{"Hello": 3, "World": 2}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got '%#v' want '%#v'", got, want)
	}
}

func TestPrint(t *testing.T) {
}
