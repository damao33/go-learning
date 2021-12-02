//go:build ignore
// +build ignore

/*
 * @Author: damao
 * @Date: 2021-12-01 18:55:15
 */

package main

import (
	"fmt"
	"learning-1/gopl.io/ch5/links"
	"log"
	"os"
)

var tokens = make(chan struct{}, 20)

//crawl2保证一个时间只能有20次调用
func crawl2(url string) []string {
	fmt.Println(url)

	tokens <- struct{}{}
	list, err := links.Extract(url)
	<-tokens

	if err != nil {
		log.Print(err)
	}
	return list
}

func main() {
	workList := make(chan []string)
	var n int
	n++
	go func() { workList <- os.Args[1:] }()

	seen := make(map[string]bool)
	for ; n > 0; n-- {
		list := <-workList
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				n++
				go func(link string) {
					workList <- crawl2(link)
				}(link)
			}
		}
	}
}
