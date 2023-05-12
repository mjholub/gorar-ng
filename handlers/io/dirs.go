package iohandlers

import (
	"fmt"
	"os"
)

type Dir struct {
	Path string
}

func (d *Dir) Create() error {
	err := os.MkdirAll(d.Path, 0o755)
	if err != nil {
		return fmt.
			Errorf("%s: making directory: %v", d.Path, err)
	}
	return nil
}
