package main

// 47.执行下面代码输出什么？
import "fmt"

func main() {
	five := []string{"Annie", "Betty", "Charley", "Doug", "Edward"}

	for _, v := range five {
		five = five[:2]
		fmt.Printf("v[%s]\n", v)
	}
}

/**
考点：**range副本机制**
循环内的切片值会缩减为2，但循环将在切片值的自身副本上进行操作。 这允许循环使用原始长度进行迭代而没有任何问题，因为后备数组仍然是完整的。
```
Output:
v[Annie]
v[Betty]
v[Charley]
v[Doug]
v[Edward]
```
**/
