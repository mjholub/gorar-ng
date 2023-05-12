package iohandlers

import (
	"bufio"
)

type DirHandler interface {
	Create() error
}

type FileHandler interface {
	WriteNew(bufRc *bufio.Reader) error
}
