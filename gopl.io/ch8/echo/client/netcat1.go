/*
 * @Author: damao
 * @Date: 2021-11-24 16:28:52
 */
package main

import (
	"io"
	"log"
	"net"
	"os"
)

func netcat1Main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8000")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	mustCopy(os.Stdout, conn)
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
