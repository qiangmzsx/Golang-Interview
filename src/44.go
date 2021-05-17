package main

// 44.以下代码有什么问题?
func Stop(stop <-chan bool) {
	close(stop)
}

/**
考点:**close channel**
有方向的channel不可能关闭

**/
