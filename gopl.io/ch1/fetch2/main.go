package main

import (
	"os"
	"fmt"
	"net/http"
	"io"
)

func main() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			continue
		}
		_, err = io.Copy(os.Stdout, resp.Body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: copying to stdout %s: %v\n", url, err)
			os.Exit(1)
		}
		defer resp.Body.Close()
	}
}
