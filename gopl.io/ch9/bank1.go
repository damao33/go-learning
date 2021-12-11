//go:build ignore
// +build ignore

/*
 * @Author: damao
 * @Date: 2021-12-07 15:12:50
 */

// Package bank provides a concurrency-safe bank with one account.
package bank

var deposits = make(chan int)
var balances = make(chan int)

func Deposit(amount int) { deposits <- amount }
func Balance() int       { return <-balances }

func teller() {
	//balance被局限在一个teller goroutine里，保证线程安全
	var balance int
	for {
		select {
		case amount := <-deposits:
			balance += amount
		case balances <- balance:
		}
	}
}

func init() {
	go teller()
}
