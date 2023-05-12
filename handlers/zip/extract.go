package ziphandlers

import (
	"archive/zip"
	"bufio"
	"fmt"
	"path/filepath"
	"strings"

	iohandlers "github.com/154pinkchairs/gorar-ng/handlers/io"

	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
)

type ZipArchive struct {
	Path     string
	files    []string
	password string
}

func (z *ZipArchive) Extract(destination *iohandlers.Dir) error {
	r, err := zip.OpenReader(z.Path)
	if err != nil {
		return err
	}
	defer r.Close()

	return unzipAll(&r.Reader, destination)
}

func unzipAll(
	r *zip.Reader,
	destination *iohandlers.Dir,
) error {
	for _, zf := range r.File {
		if err := unzipFile(zf, destination); err != nil {
			return err
		}
	}

	return nil
}

func unzipFile(
	zf *zip.File,
	destination *iohandlers.Dir,
) error {
	destFile := iohandlers.File{
		// Clean the name for security's sake (avoid Zip Slip vulnerability)
		Path: filepath.Clean(
			filepath.Join(destination.Path, zf.Name)),
		Permissions: zf.FileInfo().Mode(),
	}

	// Convert from ShiftJIS to UTF-8 (if needed)
	if strings.Contains(zf.Name, ".txt") {
		rc, err := zf.Open()
		if err != nil {
			return fmt.Errorf("%s: open compressed file: %v", zf.Name, err)
		}
		utf8Reader := transform.NewReader(rc, japanese.ShiftJIS.NewDecoder())
		bufRc := bufio.NewReader(utf8Reader)
		return destFile.WriteNew(bufRc)
	}

	if strings.HasSuffix(zf.Name, "/") {
		return destination.Create()
	}

	rc, err := zf.Open()
	if err != nil {
		return fmt.Errorf("%s: open compressed file: %v", zf.Name, err)
	}
	defer rc.Close()

	bufRc := bufio.NewReader(rc)

	return destFile.WriteNew(bufRc)
}
