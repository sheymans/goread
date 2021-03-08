package main

import (
	"flag"
	"fmt"
	"github.com/sheymans/goread/app"
	"os"
)

// TODos

// 1. add tests
// 2. do not do Println but properly log and exit
// 3. re-organize data structs (library vs books)
// 4. pretty print output
// 5. error checking parse arguments.

func main() {
	var goodReadsCSV string
	var library string
	flag.StringVar(&goodReadsCSV, "g", "", "The path to your Goodreads CSV")
	flag.StringVar(&library, "l", "smpl", "Your library")
	flag.Parse()

	err := app.Run(goodReadsCSV, library)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
