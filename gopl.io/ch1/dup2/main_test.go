package main

import (
	"bytes"
	"os"
	"reflect"
	"testing"
)

func TestCount(t *testing.T) {

	assertEquality := func(t *testing.T, got map[string]int, want map[string]int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got '%#v', want '%#v'", got, want)
		}
	}

	t.Run("when command line argument is given", func(t *testing.T) {
		// Our dependency in the program is our source of input which we will mock
		// using a buffer of bytes, hence satisfying io.Reader interface.
		b := []byte("Hello\nHello World\nHello\nWorld\nWorld\nHello\n")

		// set up our io.Reader
		// this returns a *Buffer which satisfies io.Reader
		buffer := bytes.NewBuffer(b)

		// set up our map
		counts := make(map[string]int)

		// Call the function
		Count(buffer, counts)

		want := map[string]int{"Hello": 3, "Hello World": 1, "World": 2}
		assertEquality(t, counts, want)
	})

	t.Run("when file names are given", func(t *testing.T) {
		// we can test the files by creating two temporary files with some input in it.
		filenames := []string{"test1.txt", "test2.txt"}

		// set up our map
		counts := make(map[string]int)

		for _, filename := range filenames {
			f, err := os.Open(filename)
			if err != nil {
				t.Errorf("Error opening file %v", filename)
			}
			defer f.Close()
			Count(f, counts)
		}

		want := map[string]int{"Hello": 2, "World": 2, "Hello World": 1}

		assertEquality(t, counts, want)
	})
}

func TestPrint(t *testing.T) {
	// while printing out the map we have a dependency in using the print command
	// so we will use a buffer of bytes to store the output
	buffer := bytes.Buffer{}

	// map
	counts := map[string]int{"Hello": 1, "World": 3, "Hello World": 2}

	Print(&buffer, counts)

	got := buffer.String()
	want := "3\tWorld\n2\tHello World\n"

	if got != want {
		t.Errorf("got '%#v' want '%#v'", got, want)
	}
}
