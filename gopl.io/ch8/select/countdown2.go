//go:build ignore
// +build ignore

/*
 * @Author: damao
 * @Date: 2021-12-02 12:14:59
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

	tick := time.Tick(1 * time.Second)
	countDone := make(chan int)
	go func() {
		for count := 10; count > 0; count-- {
			fmt.Println("counting: ", count)
			<-tick
		}
		countDone <- 1
	}()
	select {
	/* 	case <-time.After(10 * time.Second):
	//do nothing */
	case <-countDone:
		fmt.Println("countdown finished")
	case <-abort:
		fmt.Println("launch abort, return")
		return
	}
	fmt.Println("launch")
}
