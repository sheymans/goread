package book

import (
	"encoding/csv"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
)

type Book struct {
	title string
	author string
	isbn string
}

func (b Book) LibrarySearchUrl(library string) string {
	base := "http://" + library + ".bibliocommons.com/v2/search?query"
	if b.isbn != "=\"\"" {
		return base + b.isbn
	}
	return base + "=\"" + url.QueryEscape(b.title) + "\""
}

func (b Book) Pretty(library string) {
	fmt.Print(b.title, " : ", b.LibrarySearchUrl(library), " : ...")
	isAvailable := b.IsAvailable(library)
	fmt.Println(isAvailable)
}

func GetBooks(csvFile string) ([]Book, error) {
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

func (b Book) IsAvailable(library string) bool {
	url := b.LibrarySearchUrl(library)
	response, err := http.Get(url)
	if err != nil {
		fmt.Println("could not get URL", url, " with error ", err)
		return false
	}

	defer response.Body.Close()

	html, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("could not read HTML", err)
		return false
	}

	return strings.Contains(string(html), "availability-status-available")
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
