/*
 * @Author: damao
 * @Date: 2021-11-25 14:55:27
 */
package main

import (
	"fmt"
)

//out是只用于输出的单向channel
func counter(out chan<- int) {
	for x := 0; x < 5; x++ {
		out <- x
	}
	close(out)
}

func square(out chan<- int, in <-chan int) {
	for x := range in {
		out <- x * x
	}
	close(out)
}

func print(in <-chan int) {
	for x := range in {
		fmt.Println(x)
	}
}

func main() {
	natruals := make(chan int)
	squares := make(chan int)
	go counter(natruals)
	go square(squares, natruals)
	print(squares)
}
