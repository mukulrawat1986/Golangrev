package main_test

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/mukulrawat1986/Golangrev/Temp"
	"bytes"
)

func setup() {
	main.ActorNames = []string{}
}

func Test_AskForName(t *testing.T) {
	setup()

	a := assert.New(t)
	b := []byte("Brad\n")

	r := bytes.NewBuffer(b)
	main.AskForName(r)

	a.Equal(len(main.ActorNames), 1)
	a.Equal(main.ActorNames[0], "Brad")
}

func Test_AskForNames(t *testing.T) {
	setup()

	a := assert.New(t)
	b := []byte("Brad\nTom\nn\n")

	r := bytes.NewBuffer(b)
	main.AskForNames(r)

	a.Equal(len(main.ActorNames), 2)
	a.Equal(main.ActorNames[0], "Brad")
	a.Equal(main.ActorNames[1], "Tom")
}

func Test_AskForNames_FourNames(t *testing.T) {
	setup()

	a := assert.New(t)
	b := []byte("Brad\nTom\ny\nMary\ny\nSmith\nn\n")

	r := bytes.NewBuffer(b)
	main.AskForNames(r)

	a.Equal(len(main.ActorNames), 4)
	a.Equal(main.ActorNames[0], "Brad")
	a.Equal(main.ActorNames[1], "Tom")
	a.Equal(main.ActorNames[2], "Mary")
	a.Equal(main.ActorNames[3], "Smith")
}