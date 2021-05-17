package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	runtime.GOMAXPROCS(1)
	wg := sync.WaitGroup{}
	wg.Add(20)
	// 全部打印10
	for i := 0; i < 10; i++ {
		go func() {
			fmt.Println("i-: ", i)
			wg.Done()
		}()
	}
	// 0~9，区别在于闭包
	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Println("i+: ", i)
			wg.Done()
		}(i)
	}
	wg.Wait()
}
