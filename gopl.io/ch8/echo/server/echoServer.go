/*
 * @Author: damao
 * @Date: 2021-11-25 10:14:50
 */
package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

func echoMain() {
	listener, err := net.Listen("tcp", "127.0.0.1:8000")
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
		}
		go echoHandleConn(conn)
	}
}

func main() {
	echoMain()
}

func echoHandleConn(conn net.Conn) {
	input := bufio.NewScanner(conn)
	for input.Scan() {
		echo(conn, input.Text(), 1*time.Second)
	}
	conn.Close()
}

func echo(conn net.Conn, s string, delay time.Duration) {
	fmt.Fprintln(conn, "\t", strings.ToUpper(s))
	time.Sleep(delay)
	fmt.Fprintln(conn, "\t", s)
	time.Sleep(delay)
	fmt.Fprintln(conn, "\t", strings.ToLower(s))
	time.Sleep(delay)
}
