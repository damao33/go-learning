//go:build ignore
// +build ignore

/*
 * @Author: damao
 * @Date: 2021-12-11 13:26:35
 */

package memo

// A Memo caches the results of calling a Func.
type Memo struct {
	f     Func
	cache map[string]result
}

type Func func(key string) (interface{}, error)

type result struct {
	value interface{}
	err   error
}

func New(f Func) *Memo {
	return &Memo{f, make(map[string]result)}
}

//并发不安全
func (memo *Memo) Get(key string) (interface{}, error) {
	res, ok := memo.cache[key]
	if !ok {
		res.value, res.err = memo.f(key)
		memo.cache[key] = res
	}
	return res.value, res.err
}
