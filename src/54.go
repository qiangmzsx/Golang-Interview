package main

import (
	"math/rand"
	"time"
)

// 54.下面代码有什么问题？
func main() {
	rank := rand.New(rand.NewSource(time.Now().UnixNano()))
	//rand.
	for i := 0; i < 15; i++ {
		go func(i int) {
			for {
				_ = rank.Int()
			}
		}(i)
	}
	time.Sleep(10 * time.Second)
}

/**
考点：**rand并发**
运行会出现panic
```
panic: runtime error: index out of range

goroutine 37 [running]:
math/rand.(*rngSource).Int63(0xc42006f500, 0x62d95ae8465a29ca)
	/Users/qiangmzsx/project/go_1.8.3/src/math/rand/rng.go:231 +0x8c
math/rand.(*Rand).Int63(0xc4200740c0, 0x62d95ae8465a29ca)
	/Users/qiangmzsx/project/go_1.8.3/src/math/rand/rand.go:81 +0x33
math/rand.(*Rand).Int(0xc4200740c0, 0x62d95ae8465a29ca)
	/Users/qiangmzsx/project/go_1.8.3/src/math/rand/rand.go:99 +0x2b
```
这是因为并发时候rand会出现并发读写导致`index out of range`。
改为
```go
func main() {
	for i := 0; i < 15; i++ {
		go func(i int) {
			for {
				_ = rand.Int()
			}
		}(i)
	}
	time.Sleep(10 * time.Second)
}
```

**/
