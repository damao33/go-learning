/*
 * @Author: damao
 * @Date: 2021-11-25 10:45:09
 */
package main

import (
	"log"
	"net"
	"os"
)

func netcat2Main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8000")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	go mustCopy(os.Stdout, conn)
	mustCopy(conn, os.Stdin)
}

func main() {
	netcat2Main()
}
