package main

import (
	"fmt"
	"sync"
)

//下面的迭代会有什么问题？

type threadSafeSet struct {
	sync.RWMutex
	s []interface{}
}

func (set *threadSafeSet) Iter() <-chan interface{} {
	ch := make(chan interface{})
	go func() {
		set.RLock()

		for elem := range set.s {
			ch <- elem
			println(elem)
		}

		close(ch)
		set.RUnlock()

	}()
	return ch
}

func main() {

	th := threadSafeSet{
		s: []interface{}{"1", "2"},
	}
	v := <-th.Iter()
	fmt.Sprintf("%s%v", "ch", v)
}

/*
func (set *threadSafeSet) Iter() <-chan interface{} {
    ch := make(chan interface{})
    go func() {
        set.RLock()

        for elem := range set.s {
            ch <- elem
        }

        close(ch)
        set.RUnlock()

    }()
    return ch
}

看到这道题，我也在猜想出题者的意图在哪里。
chan?sync.RWMutex?go?迭代?
所以只能再读一次题目，就从迭代入手看看。
既然是迭代就会要求set.s全部可以遍历一次。但是chan是为缓存的，那就代表这写入一次就会阻塞。
我们把代码恢复为可以运行的方式，看看效果

```
type threadSafeSet struct {
    sync.RWMutex
    s []interface{}
}

func (set *threadSafeSet) Iter() <-chan interface{} {
    ch := make(chan interface{})
    go func() {
        set.RLock()

        for elem := range set.s {
            ch <- elem
            println(elem)
        }

        close(ch)
        set.RUnlock()

    }()
    return ch
}

func main()  {

    th:=threadSafeSet{
        s:[]interface{}{"1","2"},
    }
    v:=<-th.Iter()
    println("ch:",v)
}
```


*/
