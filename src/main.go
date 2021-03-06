package main

import (
	"flag"
	"fmt"
	"goread/m/book"
	"time"
)

func main() {
	var goodReadsCSV string
	var library string
	flag.StringVar(&goodReadsCSV, "g", "", "The path to your Goodreads CSV")
	flag.StringVar(&library, "l", "smpl", "Your library")
	flag.Parse()

	books, err := book.GetBooks(goodReadsCSV)
	if err != nil {
		fmt.Println("something went wrong: ", err)
		return
	}

	for _, book := range books {
		book.Pretty(library)
		// don't be annoying:
		time.Sleep(5 * time.Second)
	}

}
