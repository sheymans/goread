package parser

import (
	"encoding/csv"
	"github.com/sheymans/goread/book"
	"io"
	"os"
)

func ParseBooks(csvFile string) ([]book.Book, error) {
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

func readBooks(r *csv.Reader) []book.Book {
	books := make([]book.Book, 0)
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}

		bookshelf := record[16]
		if bookshelf == "to-read" {
			b := book.Book{
				Title:  record[1],
				Author: record[2],
				Isbn:   record[5],
			}

			books = append(books, b)
		}
	}

	return books
}

