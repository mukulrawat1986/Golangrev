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

func main() {
	// 	urls := []string{"http://www.google.com/", "http://www.yahoo.com",
	//		"http://www.bing.com", "http://bbc.co.uk"}

	urlCh := make(chan string)
	sizeCh := make(chan string)

	go worker(urlCh, sizeCh)

	urlCh <- "http://www.oreilly.com/"

	fmt.Printf("%s\n", <-sizeCh)
}
