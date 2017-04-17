package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get("http://api.theysaidso.com/qod.json")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error during Get Request: %v\n", err)
		os.Exit(1)
	}
	defer resp.Body.Close()
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error during reading content: %v\n", err)
		os.Exit(1)
	}
	fmt.Println(string(content))
}
