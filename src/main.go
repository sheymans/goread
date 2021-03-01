package main

import (
	"flag"
	"fmt"
)

func main() {
	var goodReadsCSV string
	flag.StringVar(&goodReadsCSV, "g", "", "The path to your Goodreads CSV")

	flag.Parse()

	fmt.Println(goodReadsCSV)

}
