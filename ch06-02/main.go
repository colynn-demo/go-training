package main

import (
	"fmt"
	"os"

	"github.com/colynn-demo/go-training/ch06-02/filesystem"
	"github.com/colynn-demo/go-training/ch06-02/filesystem/fat32"
)

func main() {
	// fileSytem := ntfs.NewNtfs()

	// fileSytem.Find()
	// fileSytem.Open()

	fileSytem := fat32.NewFat32()
	fileSytem.Find()
	fmt.Printf("%v\n", fileSytem.Open())

	// 接口不能被实例化的，
	var xx filesystem.FileSystem
	xx = fileSytem
	fmt.Print(xx.Open())

	// defer 的使用
	f1, err := os.Open("addd.txt")
	defer f1.Close()

	if err != nil {
		//
	}

	var p Person = &Student{
		name: "Tom",
		age:  18,
	}

	stu := p.(*Student) // 接口转为实例

}

// GetSystem ..
// interface 可以接收任何的数据类型
func GetSystem() interface{} {
	// return 0
	// return false

	// return ""
	// return 111.03
	return []string{"hello", "world"}
}
