package main

import (
	"bytes"
	"reflect"
	"testing"
)

func TestCount(t *testing.T) {
	// our dependency is the source of input which we will mock using a
	// buffer of bytes, hence satisfying the io.Reader interface
	b := []byte("Hello\nHello\nHello World\nWorld\n")

	// set up out io.Reader
	// this returns a *Buffer which satisfies io.Reader
	r := bytes.NewBuffer(b)

	// set up our map
	counts := make(map[string]int)

	// call the function
	Count(counts, r)

	got := counts
	want := map[string]int{"Hello": 2, "Hello World": 1, "World": 1}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got '%#v' want '%#v'", got, want)
	}
}
