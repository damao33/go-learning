//go:build ignore
// +build ignore

/*
 * @Author: damao
 * @Date: 2021-12-01 21:42:17
 */

package main

import "fmt"

// func main() {
// 	chnl := make(chan string, 10)
// 	chnl <- "a"
// 	chnl <- "b"
// 	chnl <- "c"
// 	chnl <- "d"
// 	chnl <- "e"
// 	close(chnl)
// 	flag := make(chan int, 1)
// 	go func() {
// 		for s := range chnl {
// 			fmt.Println("1", s)
// 		}
// 		// flag <- 1
// 		fmt.Println("flag1done")
// 	}()
// 	// <-flag
// 	go func() {
// 		for s := range chnl {
// 			fmt.Println("2", s)
// 		}
// 		flag <- 1
// 		fmt.Println("flag2done")
// 	}()
// 	<-flag
// 	fmt.Println("all done")
// }

/* func main() {
	slice := make(chan int, 10)
	flag := make(chan int)
	// done := make(chan int)

	var n sync.WaitGroup
	n.Add(4)
	go func() {
		slice <- 1
		slice <- 2
		slice <- 3
		slice <- 4
	}()

	start := time.Now()

	go func() {
		for s := range flag {
			n.Add(1)
			fmt.Println("a", s)
			slice <- s * 10
			n.Done()
		}
		fmt.Println("done1")
	}()
	go func() {
		defer n.Done()
		for s := range flag {
			n.Add(1)
			fmt.Println("b", s)
			slice <- s * 10
			n.Done()
		}
		fmt.Println("done2")
	}()
	go func() {
		n.Wait()
		close(slice)
	}()

	for n := range slice {
		flag <- n
	}

	fmt.Printf("all done, cost time: %v\n", time.Since(start).Seconds())

}
*/
/* func main() {
	chnl := make(chan int)
	close(chnl)
	fmt.Println(<-chnl)
	fmt.Println(<-chnl)
	fmt.Println(<-chnl)
}
*/
func main() {
	defer func() {
		fmt.Println("defer1")
	}()
	fmt.Println("123")
	defer func() {
		fmt.Println("defer2")
	}()
}
