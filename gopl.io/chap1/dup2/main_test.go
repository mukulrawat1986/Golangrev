package main

import (
	"bytes"
	"reflect"
	"strconv"
	"strings"
	"testing"
)

// makeMap converts a string to a map of map[string]int
func makeMap(input string) map[string]int {
	result := make(map[string]int)
	in := strings.Split(input, "\n")

	for _, str := range in {
		if str != "" {
			temp := strings.Split(str, "\t")
			val, _ := strconv.Atoi(temp[1])
			result[temp[0]] = val
		}
		return result
	}
}

func TestDup(t *testing.T) {
	// create a buffer of bytes.Buffer to store the output
	outputBuffer := &bytes.Buffer{}

	// create an input string
	input := "Hello\nHello\nHello World\nHell\nHello World\n"

	// create an input buffer from the input string
	inputBuffer := bytes.NewBufferString(input)

	// call the Dup function
	Dup(inputBuffer, outputBuffer)

	got := makeMap(outputBuffer.String())
	want := map[string]int{
		"Hello":       3,
		"Hello World": 2,
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %#v want %#v", got, want)
	}
}
