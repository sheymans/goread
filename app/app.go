package app

import (
	"github.com/sheymans/goread/book"
	"time"
)

func Run(goodReadsCSV string, library string) error {
	books, err := book.GetBooks(goodReadsCSV)
	if err != nil {
		// essentially rethrowing this
		return err
	}

	for _, book := range books {
		book.Pretty(library)
		// don't be annoying:
		time.Sleep(5 * time.Second)
	}

	// no errors
	return nil
}
