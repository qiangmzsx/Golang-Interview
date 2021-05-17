package main

// 34,执行下面的代码发生什么？
import "fmt"

type ConfigOne struct {
	Daemon string
}

func (c *ConfigOne) String() string {
	return fmt.Sprintf("print: %v", p)
}

func main() {
	c := &ConfigOne{}
	c.String()
}

/**
考点:**fmt.Sprintf**
如果类型实现String()，％v和％v格式将使用String()的值。因此，对该类型的String()函数内的类型使用％v会导致无限递归。
编译报错：
```
runtime: goroutine stack exceeds 1000000000-byte limit
fatal error: stack overflow
```
**/
