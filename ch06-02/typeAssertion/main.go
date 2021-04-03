package main

import "fmt"

// Person ..
type Person interface {
}

// Student ..
type Student struct {
	name string
	age  int
}

// Teacher ..
type Teacher struct {
	name string
	age  int
}

func main() {
	var p Person = &Teacher{
		name: "Tom",
		age:  18,
	}

	stu, ok := p.(*Student) // 接口转为实例
	if ok {
		fmt.Println(stu.age)
	} else {
		fmt.Println("panic...")
	}

}
