/*
 * @Author: damao
 * @Date: 2021-11-03 16:29:28
 */
package main

import "fmt"

func mapTest() {
	age := make(map[string]int)
	age["0"] = 2
	age["b"] = 3
	age["a"] = 4
	for k, v := range age {
		fmt.Println(k, v)
	}
}
