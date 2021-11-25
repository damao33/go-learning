/*
 * @Author: damao
 * @Date: 2021-11-16 16:32:31
 */

package ch7

import (
	"fmt"
	"sort"
	"testing"
)

func Test_sort(t *testing.T) {
	fmt.Println("before sort:")
	printBooks(Books)

	sort.Sort(byYear(Books))
	fmt.Println("sort:")
	printBooks(Books)

	sort.Sort(sort.Reverse(byYear(Books)))
	fmt.Println("reverse sort:")
	printBooks(Books)

	sort.Sort(customSort{Books, func(x, y *Book) bool {
		return x.Name < y.Name
	}})
	fmt.Println("custom sort:")
	printBooks(Books)
}

func Test_http1(t *testing.T) {
	http1Main()
}

func Test_http2(t *testing.T) {
	http2Main()
}

func Test_http3(t *testing.T) {
	http3Main()
}

func Test_http4(t *testing.T) {
	http4Main()
}

func Test_xmlSelect(t *testing.T) {
	xmlSelectMain()
}
