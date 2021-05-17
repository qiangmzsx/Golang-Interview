package main

import (
    "errors"
    "fmt"
)

var ErrDidNotWork = errors.New("did not work")

func DoTheThing(reallyDoIt bool) (err error) {
    if reallyDoIt {
        result, err := tryTheThing()
        if err != nil || result != "it worked" {
            err = ErrDidNotWork
        }
    }
    return err
}

func tryTheThing() (string,error)  {
    return "",ErrDidNotWork
}

func main() {
    fmt.Println(DoTheThing(true))
    fmt.Println(DoTheThing(false))
}

/**


因为 if 语句块内的 err 变量会遮罩函数作用域内的 err 变量，结果：
```
<nil>
<nil>
```
改为：
```
func DoTheThing(reallyDoIt bool) (err error) {
    var result string
    if reallyDoIt {
        result, err = tryTheThing()
        if err != nil || result != "it worked" {
            err = ErrDidNotWork
        }
    }
    return err
}
```

**/