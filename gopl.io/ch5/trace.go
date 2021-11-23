/*
 * @Author: damao
 * @Date: 2021-11-15 10:04:15
 */
package ch5

import (
	"fmt"
	"log"
	"time"
)

func trace(msg string) func() {
	start := time.Now()
	log.Printf("enter msg: %s at time:%s", msg, start)
	return func() {
		log.Printf("exit %s (%s)", msg, time.Since(start))
	}
}

func bigSlowOperation() {
	defer trace("bigSlowOperation")()
	// ..lots of work
	fmt.Println("bigSlowOperation: doing lots of work here..")
	time.Sleep(5 * time.Second)
}
