// user defined types
package main

import "net/http"
import "io/ioutil"
import "fmt"

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

func main() {
	w := &webPage{url: "http://www.oreilly.com"}
	w.get()
	fmt.Printf("URL: %s Error: %s Length %d", w.url, w.err, len(w.body))
}
