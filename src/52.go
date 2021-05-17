package main

import (
	"fmt"
	"unsafe"
)

//https://studygolang.com/topics/4915#commentForm
func main() {
	fun1()
	fun2()
	fun3()

	fun11()
	fun21()
	fun31()
}

func fun1() {
	a := 2
	c := (*string)(unsafe.Pointer(&a))
	*c = "44"
	fmt.Println("func1", &a, c, *c)
}
func fun2() {
	a := "654"
	c := (*string)(unsafe.Pointer(&a))
	*c = "44"
	fmt.Println("func2", &a, c, *c)
}
func fun3() {
	a := 3
	c := *(*string)(unsafe.Pointer(&a))
	c = "445"
	fmt.Println("func3", &a, c)
}

func fun11() {
	a := 2
	c := (*string)(unsafe.Pointer(&a))
	*c = "44"
	fmt.Println(*c)             // 无输出
	fmt.Println("fun11", c, *c) //fun11 0xc420070220 44
	fmt.Println("fun11", *c)    // 报错：unexpected fault address 0x2164692
}
func fun21() {
	a := "654"
	c := (*string)(unsafe.Pointer(&a))
	*c = "44"
	fmt.Println("fun21", *c)
}
func fun31() {
	a := 3
	c := *(*string)(unsafe.Pointer(&a))
	c = "445"
	fmt.Println("fun31", c)
}

/*
### 面试解析
```go

```

*/
