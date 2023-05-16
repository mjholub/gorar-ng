package iohandlers

import (
	"fmt"
	"os"
	"strings"
)

type Dir struct {
	Path string
}

func (d *Dir) Create() error {
	archiveSuffixes := []string{".zip", ".tar.gz", ".tar", ".rar"}
	for _, suffix := range archiveSuffixes {
		d.Path = strings.TrimSuffix(d.Path, suffix)
	}
	err := os.MkdirAll(d.Path, 0o755)
	if err != nil {
		return fmt.
			Errorf("%s: making directory: %v", d.Path, err)
	}
	return nil
}
