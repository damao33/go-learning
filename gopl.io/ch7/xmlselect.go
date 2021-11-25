/*
 * @Author: damao
 * @Date: 2021-11-23 17:06:26
 */
package ch7

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"strings"
)

func xmlSelectMain() {
	dec := xml.NewDecoder(os.Stdin)
	var stack []string //stack存放元素名称
	for {
		tok, err := dec.Token()
		if err == io.EOF {
			fmt.Println("error: io#EOF")
			break
		} else if err != nil {
			fmt.Fprintf(os.Stderr, "xmlselect :%v\n", err)
			return
		}
		switch tok := tok.(type) {
		case xml.StartElement:
			stack = append(stack, tok.Name.Local) //push
		case xml.EndElement:
			stack = stack[:len(stack)-1] //pop
		case xml.CharData:
			if containsAll(stack, os.Args[1:]) {
				fmt.Printf("%s: %s\n", strings.Join(stack, " "), tok)
			}
		}
	}
}

func containsAll(x, y []string) bool {
	for len(y) < len(x) {
		if len(y) == 0 {
			return true
		}
		if x[0] == y[0] {
			y = y[1:]
		}
		x = x[1:]
	}
	return false
}
