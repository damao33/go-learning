package main

import (
	"fmt"
)

func main() {
	x := 0
	switch {
	case x > 0:
		fmt.Println(1)
	default:
		fmt.Println(0)
	case x < 0:
		fmt.Println(-1)
	}

}
