package main

//下面代码会触发异常吗？请详细说明
import (
	"fmt"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(1)
	int_chan := make(chan int, 1)
	string_chan := make(chan string, 1)
	int_chan <- 1
	string_chan <- "hello"
	select {
	case value := <-int_chan:
		fmt.Println(value)
	case value := <-string_chan:
		panic(value)
	}
}

/**
只要认真写过select的人都应该可以理解这道题
select是随机的，所以谁也不知道执行的时候是否会先执行那一个case，所以不一定会触发异常


*/
