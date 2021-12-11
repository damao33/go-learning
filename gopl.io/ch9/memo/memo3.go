//go:build ignore
// +build ignore

/*
 * @Author: damao
 * @Date: 2021-12-11 13:46:51
 */

package memo

import "sync"

type Memo struct {
	f     Func
	mu    sync.Mutex
	cache map[string]result
}

type Func func(key string) (interface{}, error)

type result struct {
	value interface{}
	err   error
}

//在这两次获取锁的中间阶段，其它goroutine可以随意使用cache
func (memo *Memo) Get(key string) (interface{}, error) {
	memo.mu.Lock()
	res, ok := memo.cache[key]
	memo.mu.Unlock()

	if !ok {
		//这里并发时可能有重复调用服务的情况
		//应该避免多余重复工作
		res.value, res.err = memo.f(key)

		memo.mu.Lock()
		memo.cache[key] = res
		memo.mu.Unlock()
	}

	return res.value, res.err
}
