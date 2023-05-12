package rar

import (
	"bufio"
	"fmt"
	"io"
	"path/filepath"

	iohandlers "github.com/154pinkchairs/gorar-ng/handlers/io"
	"github.com/nwaples/rardecode"
)

type Archive struct {
	Path     string
	files    []string
	password string
}

// Extract extracts the rar archive to the destination path
func (r *Archive) Extract(destination *iohandlers.Dir) error {
	rr, err := rardecode.OpenReader(r.Path, r.password)
	if err != nil {
		return fmt.Errorf("read: failed to create reader: %v", err)
	}
	bufRr := bufio.NewReader(rr)

	for {
		header, err := rr.Next()
		if err == io.EOF {
			break
		}

		if header.IsDir {
			d := &iohandlers.Dir{
				Path: filepath.Join(destination.Path, header.Name),
			}
			err = d.Create()
			if err != nil {
				return err
			}
			continue
		}
		d := iohandlers.Dir{
			Path: filepath.Dir(filepath.Join(destination.Path, header.Name)),
		}
		err = d.Create()
		if err != nil {
			return err
		}

		f := iohandlers.File{
			Path:        filepath.Join(destination.Path, header.Name),
			Permissions: header.Mode(),
		}

		err = f.WriteNew(bufRr)
		if err != nil {
			return err
		}

	}

	return nil
}
