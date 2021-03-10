package main

import (
	"flag"
	"fmt"
	"github.com/sheymans/goread/app"
	"os"
)

// TODos

// 5. update README

func main() {
	var goodReadsCSV string
	var library string
	flag.StringVar(&goodReadsCSV, "g", "", "The path to your Goodreads CSV")
	flag.StringVar(&library, "l", "smpl", "Your library")
	flag.Parse()

	if goodReadsCSV == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	err := app.Run(goodReadsCSV, library)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
