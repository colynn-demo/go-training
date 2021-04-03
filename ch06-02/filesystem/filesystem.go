package filesystem

// FileSystem ..
type FileSystem interface {
	Open() string
	Close() bool
	Find() []string
}

// HLFileSytem  custom filesystem
type HLFileSytem interface {
	Close1() bool
}
