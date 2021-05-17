package main

import (
	"fmt"
	"math"
)

// 【面试题】反转整数
//反转一个整数，例如：
//
//例子1: x = 123, return 321
//例子2: x = -123, return -321
//
//输入的整数要求是一个 32bit 有符号数，如果反转后溢出，则输出 0

func main() {
	fmt.Println(reverse(-112))
	fmt.Println(reverse(2223))
}

func reverse(x int) (num int) {
	for x != 0 {
		num = num*10 + x%10
		x = x / 10
	}
	// 使用 math 包中定义好的最大最小值
	if num > math.MaxInt32 || num < math.MinInt32 {
		return 0
	}
	return
}
