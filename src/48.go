package main

//for 和 for range有什么区别?
/**
考点：**for range**
1. 使用场景不同
    for可以
    - 遍历array和slice
    - 遍历key为整型递增的map
    - 遍历string
    for range可以完成所有for可以做的事情，却能做到for不能做的，包括
    - 遍历key为string类型的map并同时获取key和value
    - 遍历channel
2. 实现不同
    for可以获取到的是被循环对象的元素本身,可以对其进行修改；
    for range使用值拷贝的方式代替被遍历的元素本身，是一个值拷贝，而不是元素本身。
**/
