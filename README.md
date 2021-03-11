# goread

Check availability of [Goodreads](https://goodreads.com/) _Want To Read_ books in a public library.

## Prerequisites

- Download your [Goodreads](https://goodreads.com/) books with the [Export](https://www.goodreads.com/review/import) function. 
  I will assume you stored the file at `/home/you/goodreads.csv`.
- Find your library code. I will refer to it with `libraryCode`. Examples:
    - `smpl` in https://smpl.bibliocommons.com/  for the Santa Monica Public Library (*the default*)
    - `austin` in https://austin.bibliocommons.com/ for the Austin Public Library
    - `sfpl` in https://sfpl.bibliocommons.com/ for the San Francisco Public Library 
- Install [Go](https://golang.org/). 

## Installation

Install the `goread` binary:

```shell
go install github.com/sheymans/goread@latest
```

Do `go help install` if you want to know where `goread` is. For me, it's at `$HOME/go/bin`.

## Usage

With `/home/you/goodreads.csv` and `libraryCode`, use `goread`:

```shell
goread -g=/home/you/goodreads.csv -l=libraryCode 
```

## Develop

Clone this repository. Develop. Run:

```shell
go run main.go -g=/home/you/goodreads.csv -l=libraryCode 
```

You can run the tests from the `goread` directory, containing `main.go`:

```shell
go test ./...
```
