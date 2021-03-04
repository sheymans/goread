package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"os"
)

type Book struct {
	title string
	author string
	isbn string
}

func main() {

	// TODO error handling for CLI parameters
	var goodReadsCSV string
	flag.StringVar(&goodReadsCSV, "g", "", "The path to your Goodreads CSV")
	flag.Parse()

	// Open the CSV
	f, err := os.Open(goodReadsCSV)
	if err != nil {
		fmt.Println(err)
		return
	}

	r := csv.NewReader(f)
	// skip column headers
	_, err = r.Read()
	if err != nil {
		fmt.Println("could not read headers", err)
		return
	}

	books := make([]Book, 0)
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}

		bookshelf := record[16]
		if bookshelf == "to-read" {
			book := Book{
				title:  record[1],
				author: record[2],
				isbn:   record[5],
			}
			books = append(books, book)
		}
	}
	fmt.Println(books)
}
