package main

type student struct {
	Name string
	Age  int
}

func pase_student() {
	m := make(map[string]*student)
	stus := []student{
		{Name: "zhou", Age: 24},
		{Name: "li", Age: 23},
		{Name: "wang", Age: 22},
	}
	for _, stu := range stus {
		m[stu.Name] = &stu
	}

}
func main() {
	pase_student()
}

/**
这样的写法初学者经常会遇到的，很危险！
for-range使用的为副本，导致每一个map的v都是最后一个Studet
大家可以试试打印出来：
```
func pase_student() {
    m := make(map[string]*student)
    stus := []student{
        {Name: "zhou", Age: 24},
        {Name: "li", Age: 23},
        {Name: "wang", Age: 22},
    }
    // 错误写法
    for _, stu := range stus {
        m[stu.Name] = &stu
    }

    for k,v:=range m{
        println(k,"=>",v.Name)
    }

    // 正确
    for i:=0;i<len(stus);i++  {
        m[stus[i].Name] = &stus[i]
    }
    for k,v:=range m{
        println(k,"=>",v.Name)
    }
}
func main(){
    pase_student()
}
```

*/
