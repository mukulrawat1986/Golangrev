// fetch prints the content found at url
package main

import (
	"os"
	"fmt"
	"net/http"
	"io/ioutil"
)

func main() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			continue
		}
		b, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
		defer resp.Body.Close()
		fmt.Printf("%s", b)
	}
}
