package main

//以下代码能编译过去吗？为什么？
import (
	"fmt"
)

type People interface {
	Speak(string) string
}

type Stduent struct{}

func (stu *Stduent) Speak(think string) (talk string) {
	if think == "bitch" {
		talk = "You are a good boy"
	} else {
		talk = "hi"
	}
	return
}

func main() {
	var peo People = Stduent{}
	think := "bitch"
	fmt.Println(peo.Speak(think))
}

/**
做错了！？说明你对golang的方法集还有一些疑问。
一句话：golang的方法集仅仅影响接口实现和方法表达式转化，与通过实例或者指针调用方法无关

*/
