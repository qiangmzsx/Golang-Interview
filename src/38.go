// 38.编译并运行如下代码会发生什么？

package main

import "fmt"

type Test struct {
	Name string
}

var list map[string]Test

func main() {

	list = make(map[string]Test)
	name := Test{"xiaoming"}
	list["name"] = name
	list["name"].Name = "Hello"
	fmt.Println(list["name"])
}

/**
考点:**map**
编程报错`cannot assign to struct field list["name"].Name in map`。
因为list["name"]不是一个普通的指针值，map的value本身是不可寻址的，因为map中的值会在内存中移动，并且旧的指针地址在map改变时会变得无效。
定义的是var list map[string]Test，注意哦Test不是指针，而且map我们都知道是可以自动扩容的，那么原来的存储name的Test可能在地址A，但是如果map扩容了地址A就不是原来的Test了，所以go就不允许我们写数据。你改为var list map[string]*Test试试看。

**/
