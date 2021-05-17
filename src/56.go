package main

import "fmt"

type IdType struct {
	Id   int
	Name string
}

func main() {
	list := []IdType{}
	fmt.Println(list)
	func(list []IdType) {
		list = append(list, IdType{
			Id:   3,
			Name: "c",
		})
	}(list)
	fmt.Println(list)

	func() {
		list = append(list, IdType{
			Id:   4,
			Name: "d",
		})
	}()
	fmt.Println(list)

	addSlice(&list)
	fmt.Println(list)
}

func addSlice(list *[]IdType) {
	*list = append(*list, IdType{
		Id:   5,
		Name: "e",
	})
}

/**
考点：**传值方式**
在Go语言中，函数参数都是以复制的方式（不支持以引用的方式）传递（比较特殊的是，Go语言闭包函数对外部变量是以引用的方式使用的）
```
结果：
[]
[]
[{4 d}]
[{4 d} {5 e}]
**/
