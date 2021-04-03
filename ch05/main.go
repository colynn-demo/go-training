package main

import (
	"fmt"

	"github.com/colynn-demo/go-training/ch05/room01"
)

func main() {
	fmt.Printf("4 people can sleep: %v\n", room01.CanUseItToSleep(4))
	fmt.Printf("2 people can sleep: %v\n", room01.CanUseItToSleep(2))
	fmt.Printf("2： %v\n", room01.Sleep01(2))
	fmt.Printf("4 ：%v\n", room01.Sleep01(4))
}
