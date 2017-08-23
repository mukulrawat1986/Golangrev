package main_test

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"bytes"
	"github.com/mukulrawat1986/Golangrev/Temp"
)

func Test_E2E(t *testing.T) {
	setup()
	a := assert.New(t)

	b := []byte("Brad Pitt\nAngelina Jolie\nn\n")

	r := bytes.NewBuffer(b)
	w := &bytes.Buffer{}

	main.Run(r, w)

	res := w.String()

	a.Contains(res, "You selected the following 2 actors:")
	a.Contains(res, "Brad Pitt")
	a.Contains(res, "Angelina Jolie")
}