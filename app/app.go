package app

import (
	"fmt"
	"github.com/sheymans/goread/book"
	"github.com/sheymans/goread/library"
	"github.com/sheymans/goread/parser"
	"time"
)

func Run(goodReadsCSV string, libraryCode string) error {
	books, err := parser.ParseBooks(goodReadsCSV)
	if err != nil {
		return err
	}

	l := library.Library{
		LibraryCode: libraryCode,
	}

	availableBooks := make([]book.Book, 0)

	for i, book := range books {
		fmt.Printf("Checking (%d/%d) %s...", i + 1, len(books), book.Title)
		pause(5)

		available, err := l.IsAvailable(book)
		if err != nil || !available {
			fmt.Println("\u2718")
		} else {
			availableBooks = append(availableBooks, book)
			fmt.Println("\u2714")
		}
	}

	if len(availableBooks) == 0 {
		fmt.Println("\nThere are no available books")
	} else {
		fmt.Println("\nThe following books are available:")
		for _, availableBook := range availableBooks {
			fmt.Printf("%s at \"%s\"\n", availableBook.Title, l.SearchUrl(availableBook))
		}
	}

	return nil
}

func pause(seconds int) {
	for i := 0; i < seconds; i++ {
		time.Sleep(time.Second)
		fmt.Printf(".")
	}
}
