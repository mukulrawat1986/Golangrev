// Fetch prints the content found at a url
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func main() {
	for _, url := range os.Args[1:] {
		// create a netClient with timeout of 10 seconds
		netClient := &http.Client{
			Timeout: time.Second * 10,
		}

		resp, err := netClient.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}

		b, err := ioutil.ReadAll(resp.Body)
		defer resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("%s\n", b)
	}
}
