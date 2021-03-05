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
	var goodReadsCSV string
	var library string
	flag.StringVar(&goodReadsCSV, "g", "", "The path to your Goodreads CSV")
	flag.StringVar(&library, "l", "smpl", "Your library")
	flag.Parse()

	books, err := getBooks(goodReadsCSV)
	if err != nil {
		fmt.Println("something went wrong: ", err)
		return
	}

	for _, book := range books {
		fmt.Println(librarySearchUrl(library, book))
	}

}

func getBooks(csvFile string) ([]Book, error) {
	r, err := createCsvReader(csvFile)
	if err != nil {
		return nil, err
	}

	_, err = readHeaders(r)
	if err != nil {
		return nil, err
	}

	books := readBooks(r)
	return books, nil
}

func createCsvReader(csvFile string) (*csv.Reader, error) {
	f, err := os.Open(csvFile)
	return csv.NewReader(f), err
}

func readHeaders(r *csv.Reader) ([]string, error){
	headers, err := r.Read()
	return headers, err
}

func readBooks(r *csv.Reader) []Book {
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

	return books
}

func librarySearchUrl(library string, b Book) string {
	base := library + ".bibliocommons.com/v2/search?query"
	if b.isbn != "=\"\"" {
		return base + b.isbn
	}
	fmt.Println(b)
	return base + "=\"" + b.title + "\""
}
