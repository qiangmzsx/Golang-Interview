package main

//请写出以下输入内容
import "fmt"

func main() {
	s := make([]int, 5)
	s = append(s, 1, 2, 3)
	fmt.Println(s)
}

/**
make初始化是由默认值的哦，此处默认值为0
[0 0 0 0 0 1 2 3]
大家试试改为:
```
s := make([]int, 0)
s = append(s, 1, 2, 3)
fmt.Println(s)//[1 2 3]
```
**/
