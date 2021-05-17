package main

// 49.讲述一下WaitGroup与Context控制并发的区别？
/**
考点：**WaitGroup和Context**
在golang 1.7之前大家普遍使用WaitGroup控制goroutine的并发，典型的实例：
```
func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
        defer wg.Done()
		time.Sleep(2*time.Second)
		fmt.Println("1号睡醒")

	}()
	go func() {
        defer wg.Done()
		time.Sleep(3*time.Second)
		fmt.Println("2号睡醒")
	}()
	wg.Wait()
	fmt.Println("==over==")
}
```


**/
