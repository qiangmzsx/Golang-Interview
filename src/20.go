package main

import "fmt"

const (
	x = iota
	y
	z = "zz"
	k
	p = iota
)

func main()  {
	fmt.Println(x,y,z,k,p)
}

/***
考点：**iota**
在常量组中如不指定类型和初始化值，这与上一行非空常量右值（表达式）相同
结果:
```
0 1 zz zz 4
```
