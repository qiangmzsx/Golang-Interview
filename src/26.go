package main

import "fmt"

type T1 struct {
}
func (t T1) m1(){
    fmt.Println("T1.m1")
}
type T2 = T1
type MyStruct struct {
    T1
    T2
}
func main() {
    my:=MyStruct{}
    my.m1()
}

/**
考点：**Go 1.9 新特性 Type Alias **
是不能正常编译的,异常：
```
ambiguous selector my.m1
```
结果不限于方法，字段也也一样；也不限于type alias，type defintion也是一样的，只要有重复的方法、字段，就会有这种提示，因为不知道该选择哪个。
改为:
```
my.T1.m1()
my.T2.m1()
```
type alias的定义，本质上是一样的类型，只是起了一个别名，源类型怎么用，别名类型也怎么用，保留源类型的所有方法、字段等。
**/