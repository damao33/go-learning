/*
 * @Author: damao
 * @Date: 2021-11-08 08:51:38
 */
package main

import (
	"fmt"
	"testing"
)

func Test_ch4(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.

	}
	arr := []int{6, 5, 1, 3, 0, 4}
	TreeSort(arr)
	fmt.Println(arr)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

		})
	}
}

func Test_point(t *testing.T) {
	tr1 := new(tree)
	changePoint(tr1)
	fmt.Println(tr1)
}

func changePoint(t *tree) {
	t.value = 1
	t.left = new(tree)
	t.left.value = 2
	t.right = new(tree)
	t.right.value = 3
}
