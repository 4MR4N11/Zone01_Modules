package models

import "os"

type Flags struct {
	L      bool
	UpperR bool
	LowerR bool
	A      bool
	T      bool
}

type Path struct {
	Path       string
	OpenedPath *os.File
}

type Args struct {
	Flags Flags
	Path  []Path
}
