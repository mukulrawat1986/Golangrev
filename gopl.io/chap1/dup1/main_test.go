package main

import (
	"bytes"
	"reflect"
	"strconv"
	"strings"
	"testing"
)

// makeMap converts a string to a map
func makeMap(input string) map[string]int {
	result := make(map[string]int)
	in := strings.Split(input, "\n")

	for _, str := range in {
		if str != "" {
			temp := strings.Split(str, "\t")
			val, _ := strconv.Atoi(temp[1])
			result[temp[0]] = val
		}
	}
	return result
}

func TestDup(t *testing.T) {

	// create a buffer of bytes.Buffer where printing happens
	outputBuffer := &bytes.Buffer{}

	// create a string which will contain the input
	input := "Hello World\nHello\nHello World\nhello\nHello\nHello\n"

	// Now using the string above create an input buffer
	inputBuffer := bytes.NewBufferString(input)

	// pass the input and output buffer to the Dup function
	Dup(inputBuffer, outputBuffer)

	get := makeMap(outputBuffer.String())
	want := map[string]int{
		"Hello World": 2,
		"Hello":       3,
	}

	if !reflect.DeepEqual(get, want) {
		t.Errorf("got %#v want %#v", get, want)
	}
}
