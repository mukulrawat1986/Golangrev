// Echo1 prints its command line arguments
package main

import (
	"io"
	"os"
)

func Echo(w io.Writer, args []string) {

}

func main() {
	Echo(os.Stdout, os.Args)
}
