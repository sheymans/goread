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

type Availability struct {
	Available bool
	Type string
}

func (l Library) SearchUrl(b book.Book) string {
	base := "http://" + l.LibraryCode + ".bibliocommons.com/v2/search?query"
	if b.Isbn != "=\"\"" {
		return base + b.Isbn
	}
	return base + "=\"" + url.QueryEscape(b.Title) + "\""
}

func (l Library) IsAvailable(b book.Book) (Availability, error) {
	searchUrl := l.SearchUrl(b)
	response, err := http.Get(searchUrl)

	if err != nil {
		return Availability{}, fmt.Errorf("could not get URL %s with error %s", searchUrl, err)
	}
	defer response.Body.Close()

	html, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return Availability{false, ""}, fmt.Errorf("could not read HTML: %s", err)
	}

	available := strings.Contains(string(html), "availability-status-available")
	return Availability{available, "physical"}, nil
}
