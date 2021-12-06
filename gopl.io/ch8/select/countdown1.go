//go:build ignore
// +build ignore

/*
 * @Author: damao
 * @Date: 2021-12-02 11:58:59
 */

package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("commencing countdown")
	tick := time.Tick(1 * time.Second)
	for count := 10; count > 0; count-- {
		fmt.Println(count)
		<-tick
	}
	fmt.Println("launch")
}
