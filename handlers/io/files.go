package iohandlers

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"runtime"
)

type File struct {
	Path        string
	Permissions os.FileMode
}

func (f *File) WriteNew(bufRc *bufio.Reader) error {
	err := os.MkdirAll(filepath.Dir(f.Path), 0o755)
	if err != nil {
		return fmt.
			Errorf("Error creating directory for file %s: %v", f.Path, err)
	}

	out, err := os.Create(f.Path)
	if err != nil {
		return fmt.
			Errorf("Error wriking file %s: %v", f.Path, err)
	}
	defer out.Close()

	err = out.Chmod(f.Permissions)
	if err != nil && runtime.GOOS != "windows" {
		return fmt.
			Errorf("Error changing file mode for file %s: %v", f.Path, err)
	}

	_, err = io.Copy(out, bufRc)
	if err != nil {
		return fmt.
			Errorf("Error extracting file %s: %v", f.Path, err)
	}
	return nil
}
