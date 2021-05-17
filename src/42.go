package main

// 42,执行下面的代码发生什么？
import "fmt"

type student struct {
	Name string
}

func zhoujielun(v interface{}) {
	switch msg := v.(type) {
	case *student, student:
		msg.Name = "qq"
		fmt.Print(msg)
	}
}

/**
考点:**类型转换**
msg不属于student类型，所以没有Name字段。
。
改为：
```
s := v.(student)
s.Name = "qq"
```
**/
