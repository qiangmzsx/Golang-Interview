package main
//编译执行下面代码会出现什么
var(
    size :=1024
    max_size = size*2
)
func main()  {
    println(size,max_size)
}

/**
考点：变量简短模式
变量简短模式限制：
- 定义变量同时显式初始化
- 不能提供数据类型
- 只能在函数内部使用

结果：
```
syntax error: unexpected :=
```
**/