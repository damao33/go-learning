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
	//movieMain()
	/*root := &tree{
		value: 1,
	}*/
	var root *tree
	fn(root)
	fmt.Println(root.value)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

		})
	}
}

func fn(root *tree) {
	*root = tree{
		value: 10,
	}
}
