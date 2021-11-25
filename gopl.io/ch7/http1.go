/*
 * @Author: damao
 * @Date: 2021-11-23 13:45:40
 */
package ch7

import (
	"fmt"
	"log"
	"net/http"
)

func http1Main() {
	db := database1{"shoes": 50, "socks": 5}
	log.Fatal(http.ListenAndServe("127.0.0.1:8000", db))
}

type dollar float32

func (d dollar) String() string { return fmt.Sprintf("$%.2f", d) }

type database1 map[string]dollar

func (db database1) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	for item, price := range db {
		fmt.Fprintf(w, "%s :%s\n", item, price)
	}
}
