package main

import (
	"fmt"
	"sync"
	"time"
)

// 46.实现两个go轮流输出：A1B2C3.....Z26
func main() {
	ChannelFunc()
	Channel1Func()
	MutexFunc()
}

func ChannelFunc() {
	zimu := make(chan int, 1)
	suzi := make(chan int, 1)
	zimu <- 0
	// zimu
	go func() {
		for i := 65; i <= 90; i++ {
			<-zimu
			fmt.Printf("%v", string(rune(i)))
			suzi <- i
		}
		return
	}()

	go func() {
		for i := 1; i <= 26; i++ {
			<-suzi
			fmt.Printf("%v", i)
			zimu <- i
		}
		return
	}()

	time.Sleep(1 * time.Second)
	fmt.Println()
}

func Channel1Func() {
	zimu := make(chan int)
	suzi := make(chan int)

	// zimu
	go func() {
		for i := 65; i <= 90; i++ {
			fmt.Printf("%v", string(rune(i)))
			zimu <- i
			<-suzi
		}
		return
	}()

	go func() {
		for i := 1; i <= 26; i++ {
			<-zimu
			fmt.Printf("%v", i)
			suzi <- i
		}
		return
	}()

	time.Sleep(10 * time.Second)
	fmt.Println()
}

func MutexFunc() {
	zm := sync.Mutex{}
	sz := sync.Mutex{}
	sz.Lock()
	go func() {
		for i := 65; i <= 90; i++ {
			zm.Lock()
			fmt.Printf("%v", string(rune(i)))
			sz.Unlock()
		}
		return
	}()
	go func() {
		for i := 1; i <= 26; i++ {
			sz.Lock()
			fmt.Printf("%v", i)
			zm.Unlock()
		}
		return
	}()

	time.Sleep(1 * time.Second)
	fmt.Println()
}

/**

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup
var wg1 sync.WaitGroup
var wg2 sync.WaitGroup

func PrintNumber() {
	for i := 1; i <= 26; i++ {
		wg2.Wait()
		fmt.Println(i)
		wg2.Add(1)
		wg1.Done()
	}
	wg.Done()
}

func PrintChar() {
	for k := 1; k <= 26; k++ {
		wg1.Wait()
		fmt.Println("A")
		wg1.Add(1)
		wg2.Done()
	}
	wg.Done()
}

func main() {
	wg.Add(2)
	wg1.Add(1)
	go PrintChar()
	go PrintNumber()
	wg.Wait()

}

*/
