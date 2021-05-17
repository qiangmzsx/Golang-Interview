package main

import "fmt"

// new slice,赋值
func main() {
	list := new([]int)
	//list:=make([]int,0)
	ii:=1
	list = append(*list, &ii)
	fmt.Println(list)
}
