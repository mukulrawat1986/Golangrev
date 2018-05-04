package main

import (
	"bytes"
	"reflect"
	"strconv"
	"strings"
	"testing"
)

func TestDup(t *testing.T) {

	// makeMap converts our string to a map
	makeMap := func(input string) map[string]int {
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

	// create a buffer of bytes.Buffer where the output is stored
	outBuffer := &bytes.Buffer{}

	// create the input string
	input := "Hello\nhello\nHello World\nhello\nHello\nhello World\n"

	// create an input buffer and initialize it with the input string
	inputBuffer := bytes.NewBufferString(input)

	Dup(inputBuffer, outBuffer)

	got := makeMap(outBuffer.String())
	want := map[string]int{
		"Hello": 2,
		"hello": 2,
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got '%#v' want '%#v'", got, want)
	}
}
