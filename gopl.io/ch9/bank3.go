//go:build ignore
// +build ignore

/*
 * @Author: damao
 * @Date: 2021-12-08 19:29:11
 */

package bank

import (
	"sync"
)

var (
	mu      sync.Mutex
	balance int
)

func Deposit(amount int) {
	mu.Lock()
	balance += amount
	mu.Unlock()
}

func Balance() int {
	mu.Lock()
	defer mu.Unlock()
	return balance
}

func deposit(amount int) {
	balance += amount
}

func Withdraw(amount int) bool {
	mu.Lock()
	defer mu.Unlock()
	deposit(-amount)
	if balance < 0 {
		deposit(amount)
		return false
	}
	return true
}
