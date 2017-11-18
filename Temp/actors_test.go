package main_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/mukulrawat1986/Golangrev/Temp"
	"net/http/httptest"
	"net/http"
	"fmt"
)

func Test_FetchActor_WithResults(t *testing.T) {
	a := assert.New(t)

	body := `{
		"page":1,
		"results":[
		{
			"id":287,
			"name":"Brad Pitt",
			"popularity":12,
			"profile_path":"/brad.jpg"
		}
		],
		"total_pages":1,
		"total_results":1
	}`

	FakeServer(body, func() {
		actor, err := main.FetchActor("Brad Pitt")
		a.NoError(err)
		a.Equal("Brad Pitt", actor.Name)
	})
}

func Test_FetchActor_WithNoResults(t *testing.T) {
	a := assert.New(t)

	body := `{
		"page":1,
		"results":[],
		"total_pages":1,
		"total_results":0
	}`

	FakeServer(body, func() {
		_, err := main.FetchActor("Brad Pitt")
		a.Error(err)
		a.Equal(err.Error(), "There are no search results for: Brad Pitt!")

	})
}

func FakeServer(b string, f func()) {
	root := main.ApiRoot

	ts := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(res, b)
	}))

	defer ts.Close()

	main.ApiRoot = ts.URL

	f()

	main.ApiRoot = root
}
