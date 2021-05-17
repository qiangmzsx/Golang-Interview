package main
//编译执行下面代码会出现什么
import "fmt"

func main()  {
    type MyInt1 int
    type MyInt2 = int
    var i int =9
    var i1 MyInt1 = i
    var i2 MyInt2 = i
    fmt.Println(i1,i2)
}

/**
考点：**Go 1.9 新特性 Type Alias **
基于一个类型创建一个新类型，称之为defintion；基于一个类型创建一个别名，称之为alias。
MyInt1为称之为defintion，虽然底层类型为int类型，但是不能直接赋值，需要强转；
MyInt2称之为alias，可以直接赋值。

结果:
```
cannot use i (type int) as type MyInt1 in assignment
```

**/