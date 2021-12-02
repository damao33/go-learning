/*
 * @Author: damao
 * @Date: 2021-11-25 13:58:52
 */
package main

import (
	"log"
	"net"
)

func netcat3Main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8000")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	go func() {

	}()
}
