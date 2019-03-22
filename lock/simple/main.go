package main

import (
	"sync"
)

// 全局变量
var counter int

var wg sync.WaitGroup
var l sync.Mutex

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			l.Lock()
			counter++
			defer func() {
				wg.Done()
				l.Unlock()
			}()
		}()
	}

	wg.Wait()
	println(counter)
}
