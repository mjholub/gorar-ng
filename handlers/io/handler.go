package iohandlers

import (
	"bufio"
	"os"
)

type DirHandler interface {
	Create() error
}

type FileHandler interface {
	WriteNew(bufRc *bufio.Reader) error
}
