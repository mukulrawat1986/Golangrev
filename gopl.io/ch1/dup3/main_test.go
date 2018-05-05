package main

import (
	"bytes"
	"reflect"
	"strconv"
	"strings"
	"testing"
)

func TestDup(t *testing.T) {

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

	// create a buffer of bytes.Buffer
	outBuffer := &bytes.Buffer{}

	// create input string
	input := "Hello\nhello\nHello World\nHello\nhello\n"

	Dup([]byte(input), outBuffer)

	got := makeMap(outBuffer.String())
	want := map[string]int{
		"Hello": 2,
		"hello": 2,
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got '%#v' want '%#v'", got, want)
	}

}
