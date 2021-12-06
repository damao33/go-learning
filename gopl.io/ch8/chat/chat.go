//go:build ignore
// +build ignore

/*
 * @Author: damao
 * @Date: 2021-12-07 00:08:16
 */

package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

/**
 * @description: main的作用是listen和accept
 * @param {*}
 * @return {*}
 */
func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	go broadcaster()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		//对每个链接都会产生新的协程
		go handleConn(conn)
	}
}

type client chan<- string

var (
	entering = make(chan client)
	leaving  = make(chan client)
	message  = make(chan string)
)

func broadcaster() {
	//存放所有连接的客户端发出channel的"资格"
	clients := make(map[client]bool)

	for {
		select {
		case msg := <-message:
			//把收到的消息广播给所有客户端的输出管道
			for cli := range clients {
				cli <- msg
			}
		case cli := <-entering:
			clients[cli] = true
		case cli := <-leaving:
			delete(clients, cli)
			close(cli)
		}
	}
}

func handleConn(conn net.Conn) {
	ch := make(chan string)
	go clientWriter(conn, ch)

	who := conn.RemoteAddr().String()
	ch <- "You are " + who
	message <- who + " has arrived"
	//把ch管道传到entering中，所以广播协程发送的消息可以通过ch写到conn中
	entering <- ch

	input := bufio.NewScanner(conn)
	for input.Scan() {
		message <- who + ": " + input.Text()
	}
	// NOTE: ignoring potential errors from input.Err()

	leaving <- ch
	message <- who + " has left"
	defer conn.Close()
}

func clientWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintf(conn, msg) //忽略网络错误
	}
}
