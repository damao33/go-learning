//go:build ignore
// +build ignore

/*
 * @Author: damao
 * @Date: 2021-12-08 19:22:37
 */

package bank

var (
	sema    = make(chan int, 1)
	balance int
)

func Deposit(amount int) {
	sema <- 1 //获取锁
	balance += amount
	<-sema //释放锁
}

func Balance() int {
	sema <- 1
	b := balance
	<-sema
	return b
}
