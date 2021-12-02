//go:build ignore
// +build ignore

/*
 * @Author: damao
 * @Date: 2021-12-01 20:46:06
 */

package main

import (
	"fmt"
	"learning-1/gopl.io/ch5/links"
	"log"
	"os"
)

func main() {
	workList := make(chan []string)
	unseenLinks := make(chan string)

	go func() { workList <- os.Args[1:] }()

	for i := 0; i < 20; i++ {
		go func() {
			for link := range unseenLinks {
				foundLink := crawl1(link)
				go func() { workList <- foundLink }()
			}
		}()
	}

	seen := make(map[string]bool)
	for list := range workList {
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				unseenLinks <- link
			}
		}
	}
}

func crawl1(url string) []string {
	fmt.Println(url)
	list, err := links.Extract(url)
	if err != nil {
		log.Print(err)
	}
	return list
}
