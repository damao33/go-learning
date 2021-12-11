//go:build ignore
// +build ignore

/*
 * @Author: damao
 * @Date: 2021-12-11 13:55:03
 */

package memo

import "sync"

type Memo struct {
	f     Func
	mu    sync.Mutex
	cache map[string]*entry
}

type Func func(key string) (interface{}, error)

type result struct {
	value interface{}
	err   error
}

type entry struct {
	res   result
	ready chan int //close when res is ready
}

func New(f Func) *Memo {
	return &Memo{f: f, cache: make(map[string]*entry)}
}

func (memo *Memo) Get(key string) (interface{}, error) {
	memo.mu.Lock()
	e := memo.cache[key]
	if e == nil {
		//对key的第一次请求，该goroutine负责计算值和广播ready
		e = &entry{ready: make(chan int)}
		memo.cache[key] = e
		memo.mu.Unlock()

		e.res.value, e.res.err = memo.f(key)

		close(e.ready)
	} else {
		//对key的重复请求
		memo.mu.Unlock()
		<-e.ready
	}
	return e.res.value, e.res.err
}
