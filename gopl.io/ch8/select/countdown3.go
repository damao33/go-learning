//go:build ignore
// +build ignore

/*
 * @Author: damao
 * @Date: 2021-12-02 12:35:41
 */

package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	abort := make(chan int)
	go func() {
		os.Stdin.Read(make([]byte, 1))
		abort <- 1
	}()

	/* 主程序返回后，ticker协程还在往tick中写数据，造成goroutine泄露*/
	tick := time.Tick(1 * time.Second)
	for count := 10; count > 0; count-- {
		fmt.Println("counting: ", count)
		select {
		case <-tick:
			// do nothing
		case <-abort:
			fmt.Println("launch abort!")
			return
		}
	}
	fmt.Println("launch")
}
