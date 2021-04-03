package room01

import (
	"fmt"
)

// CanUseItToSleep  ..
func CanUseItToSleep(pepleNum int) bool {
	if pepleNum > 3 {
		return false
	}
	return true
}

// Add01 ..
func Add01(arg1, arg2 int64) int64 {
	sum := arg1 + arg2
	return sum
}

// Sleep01 ..
func Sleep01(pepleNum int) bool {
	if pepleNum <= 3 {
		fmt.Print("small than 3\n")
	} else {
		fmt.Print("bit than 3\n")
	}

	return true
}
