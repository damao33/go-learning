/*
 * @Author: damao
 * @Date: 2021-11-16 16:32:31
 */
package ch7

import (
	"flag"
	"fmt"
	"strconv"
	"time"
)

var period = flag.Duration("period", 1*time.Second, "sleep period")

func sleepMain() {
	flag.Parse()
	fmt.Printf("sleep for %v", *period)
	time.Sleep(*period)
	fmt.Println()
}

type point struct {
	v int
}

func (p *point) String() string {
	return fmt.Sprintf("point (v: %d)", p.v)
}

func (p *point) Set(s string) error {
	v, err := strconv.Atoi(s)
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if v < 0 {
		return fmt.Errorf("new v:%d < 0", v)
	}
	(*p).v = v
	return nil
}

var p point

func init() {
	flag.Var(&p, "testp", "test flag#Value")
}

func pointMain() {
	fmt.Println(p)

	// flag.Parse()
	fmt.Println(p)
}
