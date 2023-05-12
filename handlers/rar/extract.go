package rar

import (
	"bufio"
	"fmt"
	"io"
	"path/filepath"

	"github.com/nwaples/rardecode"
)

type RarArchive struct {
	path     string
	files    []string
	password string
}

// Extract extracts the rar archive to the destination path
func (r *RarArchive) Extract(destination string) error {
	rr, err := rardecode.OpenReader(r.path, r.password)
	if err != nil {
		return fmt.Errorf("read: failed to create reader: %v", err)
	}
	bufRr := bufio.NewReader(rr)

	// sum := 1
	for {
		// sum += sum
		header, err := rr.Next()
		if err == io.EOF {
			break
		}

		if header.IsDir {
			err = Mkdir(filepath.Join(destination, header.Name))
			if err != nil {
				return err
			}
			continue
		}
		err = Mkdir(filepath.Dir(filepath.Join(destination, header.Name)))
		if err != nil {
			return err
		}

		err = WriteNewFile(filepath.Join(destination, header.Name), bufRr, header.Mode())
		if err != nil {
			return err
		}

	}

	return nil
}
