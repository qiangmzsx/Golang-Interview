package main

import (
	"fmt"
	"unsafe"
)

type Config struct {
	flag struct{}
	Name string
	Size int
}

func main() {
	conf := Config{struct{}{}, "haha", 1} // doesn't compile
	// _ = Config{Name: "bar", Size: 123} // compile okay
	// spew.Dump(conf)
	fmt.Println(conf)
	fmt.Println("size:", unsafe.Sizeof(conf.flag))
}

/**
考点：**非导出字段**
非导出字段，且为零，不能赋值。go 1.10后修复


**/
