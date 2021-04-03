package ntfs

import (
	"github.com/colynn-demo/go-training/ch06-02/filesystem"
)

// Ntfs ..
type Ntfs struct {
}

// NewNtfs .
func NewNtfs() filesystem.FileSystem {
	return &Ntfs{}
}

// Open ..
func (n *Ntfs) Open() string {
	return ""
}

// Close ..
func (n *Ntfs) Close() bool {
	return true

}

// Find ..
func (n *Ntfs) Find() []string {
	return []string{}
}
