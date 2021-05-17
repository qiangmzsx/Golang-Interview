package main

// 下面的代码有什么问题?
import (
	"fmt"
	"sync"
	"time"
)

type UserAges struct {
	ages map[string]int
	sync.Mutex
}

func (ua *UserAges) Add(name string, age int) {
	ua.Lock()
	defer ua.Unlock()
	ua.ages[name] = age
}

func (ua *UserAges) Get(name string) int {
	if age, ok := ua.ages[name]; ok {
		return age
	}
	return -1
}

func main() {
	user := UserAges{
		ages: map[string]int{
			"a": 1,
			"b": 2,
			"c": 3,
		},
	}
	go func() {
		for i := 0; i < 10; i++ {
			user.Add("aa"+fmt.Sprintf("%v", i), i*7)
			println("add:", "aa"+fmt.Sprintf("%v", i), i*7)
		}
	}()

	go func() {
		for i := 0; i < 10; i++ {
			user.Get("aa" + fmt.Sprintf("%v", i))
			println("get:", "aa"+fmt.Sprintf("%v", i))
		}
	}()
	time.Sleep(5 * time.Second)
}

/**
很不幸，会出现`fatal error: concurrent map read and map write`

修改一下看看效果
```
func (ua *UserAges) Get(name string) int {
    ua.Lock()
    defer ua.Unlock()
    if age, ok := ua.ages[name]; ok {
        return age
    }
    return -1
}
```
*/
