package main

import (
	"fmt"
	"math"
)

// 45.实现一个函数可以根据指定的size切割切片为多个小切片

func main() {
	lenth := 11
	size := 5
	list := make([]int, 0, lenth)
	for i := 0; i < lenth; i++ {
		list = append(list, i)
	}
	SpiltList(list, size)
}

func SpiltList(list []int, size int) {
	lens := len(list)
	mod := math.Ceil(float64(lens) / float64(size))
	spliltList := make([][]int, 0)
	for i := 0; i < int(mod); i++ {
		tmpList := make([]int, 0, size)
		fmt.Println("i=", i)
		if i == int(mod)-1 {
			tmpList = list[i*size:]
		} else {
			tmpList = list[i*size : i*size+size]
		}
		spliltList = append(spliltList, tmpList)
	}
	for i, sp := range spliltList {
		fmt.Println(i, " ==> ", sp)
	}
}
