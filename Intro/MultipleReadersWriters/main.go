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
	url := "http://www.google.com/"

	pageLength, err := getPage(url)
	if err != nil {
		os.Exit(1)
	}

	fmt.Printf("%s is length %d\n", url, pageLength)
}
