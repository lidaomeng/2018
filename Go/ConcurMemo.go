// Purpose: 知识点练习
// Description: 并发非阻塞缓存
// Author: lidaomeng
// Date: 2018-05-20
package main

import (
	"fmt"
	"log"
	"sync"
)

type Func func(key string) (interface{}, error)
type result struct {
	value interface{}
	err   error
}

type entry struct {
	res   result
	ready chan struct{}
}

//Version_1
func New(f Func) *Memo {
	return &Memo{f: f, cache: make(map[string]*entry)}
}

type Memo struct {
	f     Func
	mu    sync.Mutex
	cache map[string]*entry
}

func (memo *Memo) Get(key string) (value interface{}, err error) {
	memo.mu.Lock()
	e := memo.cache[key]
	if e == nil {
		e = &entry{ready: make(chan struct{})}
		memo.cache[key] = e
		memo.mu.Unlock()

		e.res.value, e.res.err = memo.f(key)

		close(e.ready)
	} else {
		memo.mu.Unlock()

		<-e.ready
	}

	return e.res.value, e.res.err
}

//// Version_2
//
//type request struct {
//	key      string
//	response chan<- result
//}
//
//type Memo struct {
//	requests chan<- request
//}
//
//func New(f Func) *Memo {
//	memo := &Memo{requests: make(chan request)}
//	go memo.server(f)
//	return memo
//}
//
//func (memo *Memo) Get(key string) (interface{}, error) {
//	response := make(chan result)
//	memo.requests <- request{key, response}
//	res := <-response
//	return res.value, res.err
//}
//
//func (memo *Memo) Close() {
//	close(memo.requests)
//}
//
//func (memo *Memo) server(f Func) {
//	cache := make(map[string]*entry)
//	for req := range memo.requests {
//		e := cache[req.key]
//		if e == nil {
//			e = &entry{ready: make(chan struct{})}
//			cache[req.key] = e
//			go e.call(f, req.key)
//		}
//		go e.deliver(req.response)
//	}
//}
//
//func (e *entry) call(f Func, key string) {
//	e.res.value, e.res.err = f(key)
//}
//
//func (e *entry) deliver(response chan<- result) {
//	<-e.ready
//	response <- e.res
//}

func myFunc(key string) (interface{}, error) {
	return "*" + key + "*", nil
}

func main() {
	keys := []string{"li", "dao", "meng", "li", "dao", "meng", "li", "dao", "meng"}

	memo := New(myFunc)

	fmt.Println("Begin")
	var n sync.WaitGroup

	for _, key := range keys {
		//fmt.Println(key)
		n.Add(1)
		go func(key string) {
			defer n.Done()
			elem, err := memo.Get(key)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(key, " -> ", elem)
		}(key)
	}

	//time.Sleep(5 * time.Second)
	n.Wait()
	fmt.Println("End")
}
