package main

// 50.编译并运行如下代码会发生什么？
import "fmt"

func main() {
	mmap := make(map[map[string]string]int, 0)
	mmap[map[string]string{"a": "a"}] = 1
	mmap[map[string]string{"b": "b"}] = 1
	mmap[map[string]string{"c": "c"}] = 1
	fmt.Println(mmap)
}

/**
考点:**map key类型**
golang中的map，的 key 可以是很多种类型，比如 bool, 数字，string, 指针, channel , 还有 只包含前面几个类型的 interface types, structs, arrays
显然，slice， map 还有 function 是不可以了，因为这几个没法用 == 来判断，即不可比较类型。
可以将map[map[string]string]int改为map[struct]int
**/
