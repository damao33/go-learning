/*
 * @Author: damao
 * @Date: 2021-11-08 10:30:13
 */

package main

type tree struct {
	value       int
	left, right *tree
}

func TreeSort(arr []int) {
	var root *tree
	for _, v := range arr {
		root = add(root, v)
	}
	appendValues(arr[:0], root)
}
func TreeSort1(arr []int) {
	var root *tree
	for _, v := range arr {
		add1(root, v)
	}
	appendValues(arr[:0], root)
}

func appendValues(arr []int, root *tree) []int {
	if root != nil {
		arr = appendValues(arr, root.left)
		arr = append(arr, root.value)
		arr = appendValues(arr, root.right)
	}
	return arr
}

func add(root *tree, v int) *tree {
	if root == nil {
		root = new(tree)
		root.value = v
		return root
	}
	if v < root.value {
		root.left = add(root.left, v)
	} else {
		root.right = add(root.right, v)
	}
	return root
}
func add1(root *tree, v int) {
	if root == nil {
		root = &tree{
			value: v,
		}
		return
	}
	if v < root.value {
		if root.left == nil {
			root.left = &tree{
				value: v,
			}
			return
		}
		add1(root.left, v)

	} else {
		if root.right == nil {
			root.right = &tree{
				value: v,
			}
			return
		}
		add1(root.right, v)

	}
}
