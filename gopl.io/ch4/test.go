package main

import (
	"fmt"
	"unicode"
)

func mainTest() {
	arr := [5]int{1, 2, 3, 4, 5}
	reverse(&arr)
	fmt.Println("reverse: ", arr)

	arr1 := [5]int{1, 2, 3, 4, 5}
	fmt.Println("roatate: ", rotate(2, arr1[:]))

	strs := []string{"a", "a", "a", "b", "b", "d", "b", "c", "c"}
	fmt.Println("nonEqualsWithNext: ", nonEqualsWithNext(strs))

	bytes := []byte("aa  b c    d")
	bytes = removeSpace(bytes)
	fmt.Println("remove space: ", string(bytes))

	bytes1 := []byte("123456")
	bytes1 = reverseBytes(bytes1)
	fmt.Println(string(bytes1))
}

/**
 * @description: 重写reverse函数，使用数组指针代替slice
 * @param {*[5]int} arr
 * @return {*}
 */
func reverse(arr *[5]int) {
	i, j := 0, len(*arr)-1
	for i < j {
		(*arr)[i], (*arr)[j] = (*arr)[j], (*arr)[i]
		i++
		j--
	}
}

/**
 * @description: 编写一个rotate函数，通过一次循环完成旋转
 * @param {int} n
 * @param {[]int} arr
 * @return {*}
 */
func rotate(n int, arr []int) []int {
	res := make([]int, len(arr))
	for i := range arr {
		res[i] = arr[(i+n)%len(arr)]
	}
	return res
}

/**
 * @description: 写一个函数在原地完成消除[]string 中相邻重复的字符串的操作
 * @param {[]string} strs
 * @return {*}
 */
func nonEqualsWithNext(strs []string) []string {
	// range迭代的i不会因为i--而改变
	/* for i := range strs {
		if i+1 >= len(strs) {
			break
		}
		if strs[i] == strs[i+1] {
			copy(strs[i:], strs[i+1:])
			strs = strs[:len(strs)-1]
			i--
		}
	} */
	for i := 0; i < len(strs)-1; i++ {
		if strs[i] == strs[i+1] {
			copy(strs[i:], strs[i+1:])
			strs = strs[:len(strs)-1]
			i--
		}
	}

	return strs
}

/**
 * @description: 编写一个函数,原地将一个 UTF-8 编码的[]byte 类型的 slice 中相邻的空格
 * @param {[]byte} bytes
 * @return {*}
 */
func removeSpace(bytes []byte) []byte {
	for i := 0; i < len(bytes); i++ {
		if unicode.IsSpace(rune(bytes[i])) && unicode.IsSpace(rune(bytes[i+1])) {
			copy(bytes[i:], bytes[i+1:])
			bytes = bytes[:len(bytes)-1]
			i--
		}
	}
	return bytes
}

/**
 * @description: 修改 reverse 函数用于原地反转 UTF-8 编码的[]byte。是否可以不用分配额外的内存
 * @param {[]byte} bytes
 * @return {*}
 */
func reverseBytes(bytes []byte) []byte {
	for i, j := 0, len(bytes)-1; i < j; i, j = i+1, j-1 {
		bytes[i], bytes[j] = bytes[j], bytes[i]
	}
	return bytes
}
