package main

import "fmt"

func main() {
	a := 1

	for i := 0; i < 5; i++ {
		a++
		if a == 2 {
			break LOOP
			//goto LOOP
		}
	}
LOOP:
	fmt.Printf("a=%v\n", a)
}

/**
考点：**break label**
编译会出现错误
```
mianshi/55.go:10:10: break label not defined: LOOP
mianshi/55.go:14:1: label LOOP defined and not used
```
break label,break的跳转标签(label)必须放在循环语句for前面，并且在break label跳出循环不再执行for循环里的代码。
break标签只能用于for循环。
如
```go
func main() {
	a := 1
LOOP:
	fmt.Printf("a=%v\n", a)
	for i := 0; i < 5; i++ {
		a++
		if a == 2 {
			break LOOP
			//goto LOOP
		}
	}

}
```
也会报错：
```
invalid break label LOOP
```
将源码改为：
```go
func main() {
	a := 1

	for i := 0; i < 5; i++ {
		a++
		if a == 2 {
			goto LOOP
		}
	}
LOOP:
	fmt.Printf("a=%v\n", a)
}
```


**/
