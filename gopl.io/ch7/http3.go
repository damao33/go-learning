/*
 * @Author: damao
 * @Date: 2021-11-23 14:29:29
 */
package ch7

import (
	"fmt"
	"log"
	"net/http"
)

func http3Main() {
	db := database3{"shoes": 50, "socks": 5}
	mux := http.NewServeMux()
	/*
		db#list不满足Handler接口，所以需要http#HandlerFunc转换
		mux.Handle("/list", http.HandlerFunc(db.list))
		mux.Handle("/price", http.HandlerFunc(db.price))
	*/
	mux.HandleFunc("/list", db.list)
	mux.HandleFunc("/price", db.price)
	log.Fatal(http.ListenAndServe("127.0.0.1:8000", mux))
}

type database3 map[string]dollar

func (db database3) list(w http.ResponseWriter, req *http.Request) {
	for item, price := range db {
		fmt.Fprintf(w, "%s : %s\n", item, price)
	}
}
func (db database3) price(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	price, ok := db[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound) //404
		fmt.Fprintf(w, "no such item: %q\n", item)
		return
	}
	fmt.Fprintf(w, "item %s : %s\n", item, price)
}
