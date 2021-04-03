package fat32

import (
	"github.com/colynn-demo/go-training/ch06-02/filesystem"
)

// Fat32 ..
type Fat32 struct {
}

// NewFat32 .
func NewFat32() filesystem.FileSystem {
	return &Fat32{}
}

// NewHLFat32 .
func NewHLFat32() filesystem.HLFileSytem {
	return &Fat32{}
}

var _ filesystem.FileSystem = (*Fat32)(nil)
var _ filesystem.HLFileSytem = (*Fat32)(nil)

// Open ..
func (n *Fat32) Open() string {
	return "open is call"
}

// Close ..
func (n *Fat32) Close() bool {
	return true

}

// Close01 ..
func (n *Fat32) Close1() bool {
	return true

}

// Find ..
func (n *Fat32) Find() []string {
	return []string{}
}
