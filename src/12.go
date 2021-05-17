package main

// 是否可以编译通过？如果通过，输出什么？
func main() {
	i := GetValue()

	switch i.(type) {
	case int:
		println("int")
	case string:
		println("string")
	case interface{}:
		println("interface")
	default:
		println("unknown")
	}

}

func GetValue() int {
	return 1
}

/**
考点：**type**

编译失败，因为type只能使用在interface

**/
