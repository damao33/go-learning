//go:build ignore
// +build ignore

/*
 * @Author: damao
 * @Date: 2021-12-10 21:31:44
 */

package main

import (
	"sync"
)

var mu sync.RWMutex
var balance int

/**
 * @description: 多读单写
 * @param {*}
 * @return {*}
 */
func Balance() int {
	mu.RLock()
	defer mu.RUnlock()
	return balance
}
