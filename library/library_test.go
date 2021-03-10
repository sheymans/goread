package library

import (
	"github.com/sheymans/goread/book"
	"testing"
)

func TestLibrary_SearchUrl(t *testing.T) {
	b := book.Book{
		Title:  "My Book",
		Author: "Moi",
		Isbn:   "=019231124242",
	}

	l := Library{LibraryCode: "austin"}

	searchUrl := l.SearchUrl(b)
	if searchUrl != "http://austin.bibliocommons.com/v2/search?query=019231124242" {
		t.Error("Expected 019231124242, got ", searchUrl)
	}
}
