package main

func main()  {

    for i:=0;i<10 ;i++  {
    loop:
        println(i)
    }
    goto loop
}

/***
考点：**goto**
goto不能跳转到其他函数或者内层代码
```
goto loop jumps into block starting at
```

*/