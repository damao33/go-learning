/*
 * @Author: damao
 * @Date: 2021-11-23 13:55:38
 */
package ch7

import (
	"fmt"
	"log"
	"net/http"
)

func http2Main() {
	db := database2{"shoes": 50, "socks": 5}
	log.Fatal(http.ListenAndServe("127.0.0.1:8000", db))
}

func (db database2) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case "/list":
		for item, price := range db {
			fmt.Fprintf(w, "%s : %s\n", item, price)
		}
	case "/price":
		item := req.URL.Query().Get("item")
		price, ok := db[item]
		if !ok {
			w.WriteHeader(http.StatusNotFound) //404
			fmt.Fprintf(w, "no such item: %q\n", item)
			return
		}
		fmt.Fprintf(w, "item %s : %s\n", item, price)
	default:
		w.WriteHeader(http.StatusNotFound) //404
		fmt.Fprintf(w, "no such page: %s\n", req.URL)
	}

}

type database2 map[string]dollar
