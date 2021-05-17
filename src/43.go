package main

// 43,执行下面的代码发生什么？
import (
	"encoding/json"
	"fmt"
)

type People struct {
	name string `json:"name"`
}

func main() {
	js := `{
		"name":"11"
	}`
	var p People
	err := json.Unmarshal([]byte(js), &p)
	if err != nil {
		fmt.Println("err: ", err)
		return
	}
	fmt.Println("people: ", p)
}

/**
考点:**结构体访问控制**
这道题坑很大，很多同学一看就以为是`p`的初始化问题，实际上是因为`name`首字母是小写，导致其他包不能访问，所以输出为空结构体。
。
改为：
```
type People struct {
	Name string `json:"name"`
}
```
**/
