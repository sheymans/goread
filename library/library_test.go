package library

import (
	"testing"

	"github.com/sheymans/goread/book"
)

func TestLibrary_SearchUrl(t *testing.T) {
	b := book.Book{
		Title:  "My Book",
		Author: "Moi",
		Isbn:   "=019231124242",
	}

	l := Library{LibraryCode: "austin"}

	searchUrl := l.SearchUrl(b)
	if searchUrl != "http://austin.bibliocommons.com/v2/search?query=\"My+Book\"" {
		t.Error("Expected query=\"My+Book\", got ", searchUrl)
	}
}
