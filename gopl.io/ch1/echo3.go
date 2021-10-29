package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main3() {
	start1 := time.Now()
	fmt.Println(strings.Join(os.Args[1:], "a"))
	fmt.Printf("join costs: %.2fs\n", time.Since(start1).Seconds())
	start2 := time.Now()
	var s string
	for _, arg := range os.Args {
		s += arg + " "
	}
	fmt.Println(s)
	fmt.Printf("+= costs: %.2fs\n", time.Since(start2).Seconds())
}
