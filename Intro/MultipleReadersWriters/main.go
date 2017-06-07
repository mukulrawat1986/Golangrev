package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func getPage(url string) (int, error) {
	resp, err := http.Get(url)
	if err != nil {
		return 0, err
	}

	// close the response Body
	defer resp.Body.Close()

	// read the body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, nil
	}

	return len(body), nil
}

func worker(urlCh chan string, sizeCh chan string) {
	for {
		url := <-urlCh
		length, err := getPage(url)
		if err == nil {
			sizeCh <- fmt.Sprintf("%s has length %d", url, length)
		} else {
			sizeCh <- fmt.Sprintf("Error getting %s: %s", url, err)
		}
	}
}

func generator(url string, urlCh chan string) {
	urlCh <- url
}

func main() {
	urls := []string{"http://www.google.com/", "http://www.yahoo.com",
		"http://www.bing.com", "http://bbc.co.uk", "http://www.oreilly.com"}

	urlCh := make(chan string)
	sizeCh := make(chan string)

	// run multiple workers
	for i := 0; i < 10; i++ {
		go worker(urlCh, sizeCh)
	}

	for _, url := range urls {
		go generator(url, urlCh)
	}

	for i := 0; i < len(urls); i++ {
		fmt.Printf("%s\n", <-sizeCh)
	}
}
