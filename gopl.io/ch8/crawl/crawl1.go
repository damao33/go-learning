//go:build ignore
// +build ignore

/*
 * @Author: damao
 * @Date: 2021-12-01 13:01:51
 */

package main

import (
	"fmt"
	"learning-1/gopl.io/ch5/links"
	"log"
	"os"
)

func crawl1(url string) []string {
	fmt.Println(url)
	list, err := links.Extract(url)
	if err != nil {
		log.Print(err)
	}
	return list
}

func main() {
	workList := make(chan []string)

	go func() {
		for i, arg := range os.Args {
			fmt.Println(i, arg)
		}
		workList <- os.Args[1:]
	}()

	seen := make(map[string]bool)
	for list := range workList {
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				go func(link string) {
					workList <- crawl1(link)
				}(link)
			}

		}
	}
}
