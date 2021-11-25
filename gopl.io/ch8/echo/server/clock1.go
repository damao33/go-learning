/*
 * @Author: damao
 * @Date: 2021-11-24 16:14:31
 */
package main

import (
	"io"
	"log"
	"net"
	"time"
)

func clock1Main() {
	listener, err := net.Listen("tcp", "127.0.0.1:8000")
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		//一个时间只能处理一个连接
		// handleConn(conn)
		//同时处理多个连接
		go handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	defer conn.Close()
	for {
		_, err := io.WriteString(conn, time.Now().Format("15:04:04\n"))
		if err != nil {
			log.Printf("client might have disconnected")
			return
		}
		time.Sleep(1 * time.Second)
	}
}
