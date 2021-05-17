package main

// 39.ABCD中哪一行存在错误？
type S struct {
}

func f(x interface{}) {
}

func g(x *interface{}) {
}

func main() {
	s := S{}
	p := &s
	f(s) //A
	g(s) //B
	f(p) //C
	g(p) //D

}

/**
考点:**interface**
看到这道题需要第一时间想到的是Golang是强类型语言，interface是所有golang类型的父类，类似Java的Object。
函数中`func f(x interface{})`的`interface{}`可以支持传入golang的任何类型，包括指针，但是函数`func g(x *interface{})`只能接受`*interface{}`.

```
var i interface{} = s
g(&i)
```
结果：
```
BD
``

**/
