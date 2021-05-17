// 49.解决下面问题：输出MutilParam= [ssss [1 2 3 4]]如何做到输出为[ssss 1 2 3 4]
package main

import "fmt"

func MutilParam(p ...interface{}) {
	fmt.Println("MutilParam=", p)
}
func main() {
	MutilParam("ssss", 1, 2, 3, 4) //[ssss 1 2 3 4]
	iis := []int{1, 2, 3, 4}
	MutilParam("ssss", iis) //输出MutilParam= [ssss [1 2 3 4]]如何做到输出为[ssss 1 2 3 4]
}

/**
考点：**函数变参**
这样的情况会在开源类 库如xorm升级版本后出现Exce函数不兼容的问题。
解决方式有两个：
```go
//方法一：interface[]
tmpParams := make([]interface{}, 0, len(iis)+1)
tmpParams = append(tmpParams, "ssss")
for _, ii := range iis {
    tmpParams = append(tmpParams, ii)
}
MutilParam(tmpParams...)
//方法二:反射
f := MutilParam
value := reflect.ValueOf(f)
pps := make([]reflect.Value, 0, len(iis)+1)
pps = append(pps, reflect.ValueOf("ssss"))
for _, ii := range iis {
    pps = append(pps, reflect.ValueOf(ii))
}
value.Call(pps)
```

**/
