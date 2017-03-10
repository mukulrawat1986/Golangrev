package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
)

func main() {

	// get the book moby dick
	res, err := http.Get("http://www.gutenberg.org/files/2701/old/moby10b.txt")
	if err != nil {
		log.Fatal(err)
	}

	// scan the page
	scanner := bufio.NewScanner(res.Body)
	defer res.Body.Close()

	// set the split function for the scanning operation
	scanner.Split(bufio.ScanWords)

	// create slice to hold counts
	buckets := make([]int, 12)

	// loop over the ScanWords
	for scanner.Scan() {
		n := HashBucket(scanner.Text(), 12)
		buckets[n]++
	}

	fmt.Println(buckets)
}

func HashBucket(word string, buckets int) int {
	letter := int(word[0])
	bucket := letter % buckets
	return bucket
}
