package library

import (
	"fmt"
	"github.com/sheymans/goread/book"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

type Library struct {
	LibraryCode string
}

func (l Library) SearchUrl(b book.Book) string {
	base := "http://" + l.LibraryCode + ".bibliocommons.com/v2/search?query"
	if b.Isbn != "=\"\"" {
		return base + b.Isbn
	}
	return base + "=\"" + url.QueryEscape(b.Title) + "\""
}

func (l Library) IsAvailable(b book.Book) (bool, error) {
	url := l.SearchUrl(b)
	response, err := http.Get(url)

	if err != nil {
		return false, fmt.Errorf("could not get URL %s with error %s", url, err)
	}
	defer response.Body.Close()

	html, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return false, fmt.Errorf("could not read HTML: %s", err)
	}

	return strings.Contains(string(html), "availability-status-available"), nil
}
