package main

import (
	"bufio"
	"fmt"
	"os"
)

func dedup(){
	input:=bufio.NewScanner(os.Stdin)
	fmt.Println("input start")
	for input.Scan(){
		line:=input.Text()
		fmt.Println(line)
	}
	fmt.Println("input end")
}