/*
 * @Author: damao
 * @Date: 2021-11-23 15:56:21
 */
package ch7

import (
	"log"
	"net/http"
)

func http4Main() {
	db := database3{"shoes": 50, "socks": 5}
	//http包定义了一个全局ServeMux实例DefaultServeMux和包级别的Handler和HandlerFunc
	http.HandleFunc("/list", db.list)
	http.HandleFunc("/price", db.price)
	//DefaultServeMux空值也可以工作
	log.Fatal(http.ListenAndServe("127.0.0.1:8000", nil))
}
