package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type webPage struct {
	url  string
	body []byte
	err  error
}

func (w *webPage) get() {
	resp, err := http.Get(w.url)
	if err != nil {
		w.err = err
		return
	}
	defer resp.Body.Close()

	w.body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		w.err = err
		return
	}
}

func (w *webPage) isOk() bool {
	return w.err == nil
}

type SummableSlice []int

func (s SummableSlice) sum() int {
	sum := 0
	for _, i := range s {
		sum += i
	}
	return sum
}

func main() {
	w := &webPage{url: "http://www.oreilly.com/"}
	w.get()
	if w.isOk() {
		fmt.Printf("URL: %s Error:Nil Length:%d\n", w.url, len(w.body))
	} else {
		fmt.Printf("Something went wrong\n")
	}

	var s SummableSlice = SummableSlice{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Printf("Sum is %d\n", s.sum())
}
