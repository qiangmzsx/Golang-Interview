package main

import (
	"fmt"
)

func main() {
	/*defer func() {
		if err := recover(); err != nil {
			fmt.Println("one=", err)
		}
	}()*/
	defer_call()
}
func defer_call() {
	/*defer func() {
		if err := recover(); err != nil {
			fmt.Println("one=", err)
		}
	}()*/
	defer func() { fmt.Println("打印前") }()
	defer func() { fmt.Println("打印中") }()
	defer func() { fmt.Println("打印后") }()

	panic("触发异常")

}

/*defer func() {
	if err := recover(); err != nil {
		fmt.Println("tow=", err)
	}
}()*/
/**
defer 是**后进先出**
panic 需要等defer 结束后才会向上传递

*/
