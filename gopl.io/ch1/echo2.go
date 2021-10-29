package main

import (
	"fmt"
	"os"
)

func main2() {
	var s, sep string
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println("asd" + s)
}
