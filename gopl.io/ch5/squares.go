/*
 * @Author: damao
 * @Date: 2021-11-11 10:13:09
 */
package ch5

import "fmt"

func square() func() int {
	var x int
	return func() int {
		x++
		return x * x
	}
}

func squareTest() {
	f := square()
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
}
