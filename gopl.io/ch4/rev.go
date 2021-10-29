package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5, 6}
	reverse(arr[:2])
	fmt.Println(arr)
	reverse(arr[2:])
	fmt.Println(arr)
	reverse(arr)
	fmt.Println(arr)
}

func reverse(n []int) {
	for i, j := 0, len(n)-1; i < j; i, j = i+1, j-1 {
		n[i], n[j] = n[j], n[i]
	}
}
