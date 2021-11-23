/*
 * @Author: damao
 * @Date: 2021-11-23 09:41:42
 */
package ch7

import (
	"fmt"
	"os"
	"text/tabwriter"
)

type Book struct {
	Name   string
	Author string
	Year   int
}

var Books = []*Book{
	{"book1", "author1", 2001},
	{"book2", "author2", 1988},
	{"book3", "author3", 2012},
}

func printBooks(books []*Book) {
	const format = "%v\t%v\t%v\t\n"
	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)
	fmt.Fprintf(tw, format, "Name", "Author", "Year")
	fmt.Fprintf(tw, format, "----", "------", "----")
	for _, b := range books {
		fmt.Fprintf(tw, format, b.Name, b.Author, b.Year)
	}
	tw.Flush()
}

type byYear []*Book

func (b byYear) Len() int           { return len(b) }
func (b byYear) Less(i, j int) bool { return b[i].Year < b[j].Year }
func (b byYear) Swap(i, j int)      { b[i], b[j] = b[j], b[i] }

type customSort struct {
	books []*Book
	less  func(x, y *Book) bool
}

func (c customSort) Len() int           { return len(c.books) }
func (c customSort) Less(i, j int) bool { return c.less(c.books[i], c.books[j]) }
func (c customSort) Swap(i, j int)      { c.books[i], c.books[j] = c.books[j], c.books[i] }
