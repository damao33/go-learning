/*
 * @Author: damao
 * @Date: 2021-11-24 15:40:31
 */
package ch8

import (
	"fmt"
	"time"
)

func spinnerMain() {
	go spinner(100 * time.Millisecond)
	const n = 80
	fibN := fib(n)
	fmt.Printf("\rFibonacci(%d) = %d\n", n, fibN)
}

func spinner(delay time.Duration) {
	for {
		for _, r := range `-\|/` {
			fmt.Printf("\r%c", r)
			time.Sleep(delay)
		}
	}
}

func fib(n int) int {
	if n < 2 {
		return n
	}
	return fib(n-1) + fib(n-2)
}
