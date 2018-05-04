package main

import (
	"bytes"
	"reflect"
	"strconv"
	"strings"
	"testing"
)

func TestDup(t *testing.T) {

	// private function to convert our string to a map.
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

	t.Run("when no file is given, input from console", func(t *testing.T) {
		// create a buffer to store output of type bytes.Buffer
		outBuffer := &bytes.Buffer{}

		// create an input string
		input := "Hello\nhello\nHello\nhello world\nHello\nhello\n"

		// create an input buffer using the input string
		inBuffer := bytes.NewBufferString(input)

		Dup(inBuffer, outBuffer, nil)

		got := makeMap(outBuffer.String())
		want := map[string]int{
			"Hello": 3,
			"hello": 2,
		}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got '%#v' want '%#v'", got, want)
		}
	})
}
