/*
 * @Author: damao
 * @Date: 2021-11-10 11:03:49
 */
package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type movie struct {
	Title string
	Year  int     `json:"release"`
	Range float32 `json:",omitempty"` //moitempty:零值时不生成
}

func movieMain() {
	movies := []movie{
		{"movieA", 1999, 0},
		{"movieB", 2000, 9.9},
	}
	data, err := json.Marshal(movies)
	if err != nil {
		log.Fatalf("json Marshal failed: %s", err)
	}
	fmt.Printf("json marshal: \n%s\n", data)
	fmt.Println("------------")
	// json#MarshalIndent  arg1:每行输出前缀 arg2:每一个层级的缩进
	data, err = json.MarshalIndent(movies, "", "    ")
	if err != nil {
		log.Fatalf("json MarshalIndent failed: %s", err)
	}
	fmt.Printf("json marshalIndent: \n%s\n", data)

	var titles []struct{ Title string }
	if err := json.Unmarshal(data, &titles); err != nil {
		log.Fatalf("json unmarshal failed: %s", err)
	}
	fmt.Println(titles)
}
