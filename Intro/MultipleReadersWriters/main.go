package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
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

func main() {
	urls := []string{"http://www.google.com/", "http://www.yahoo.com",
		"http://www.bing.com", "http://bbc.co.uk"}

	for _, url := range urls {
		pageLength, err := getPage(url)
		if err != nil {
			os.Exit(1)
		}

		fmt.Printf("%s is length %d\n", url, pageLength)
	}
}
