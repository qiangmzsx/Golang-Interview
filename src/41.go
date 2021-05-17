// 41,执行下面的代码发生什么？
package main

type Param map[string]interface{}

type Show struct {
	*Param
}

func main() {
	s := new(Show)
	s.Param["RMB"] = 10000
}

/**
考点:**map初始化**
map需要初始化后才能使用。
编译错误：invalid operation: s.Param["RMB"] (type *Param does not support indexing)
**/
