package main

import (
	"log"
	"os"

	_ "github.com/mukulrawat1986/Golangrev/GoInAction/rssfeed/matchers"
	"github.com/mukulrawat1986/Golangrev/GoInAction/rssfeed/search"
)

// init is called prior to main
func init() {
	// change the device for logging to stdout
	log.SetOutput(os.Stdout)
}

// main is the entry point for the program
func main() {
	// Perform the search for the specified terms
	search.Run("president")
}
