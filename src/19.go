package main

import "fmt"

func GetValue(m map[int]string, id int) (string, bool) {
	if _, exist := m[id]; exist {
		return "存在数据", true
	}
	return nil, false
}
func main()  {
	intmap:=map[int]string{
		1:"a",
		2:"bb",
		3:"ccc",
	}

	v,err:=GetValue(intmap,3)
	fmt.Println(v,err)
}
/*
在刚刚开始使用golang写代码的时候，经常会放错。给一个变量放回一个nil，这个通常编译的时候不会报错，但是运行是时候会报cannot use nil as type string in return argument的错误，对于nil，一般通常指针类型和interface类型可以使用这样的返回值
func Get(m map[int]string, id int) (string, bool) {
    if _, exist := m[id]; exist {
        return "存在数据", true
    }
    return nil, false
}
*/
