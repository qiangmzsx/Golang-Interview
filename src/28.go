package main

func test() []func()  {
    var funs []func()
    for i:=0;i<2 ;i++  {
        funs = append(funs, func() {
            println(&i,i)
        })
    }
    return funs
}

func main(){
    funs:=test()
    for _,f:=range funs{
        f()
    }
}

/**
考点：**闭包延迟求值**
for循环复用局部变量i，每一次放入匿名函数的应用都是想一个变量。
结果：
```
0xc042046000 2
0xc042046000 2
```
如果想不一样可以改为：
```
func test() []func()  {
    var funs []func()
    for i:=0;i<2 ;i++  {
        x:=i
        funs = append(funs, func() {
            println(&x,x)
        })
    }
    return funs
}
```

**/
