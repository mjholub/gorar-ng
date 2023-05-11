package main

import (
	"archive/zip"
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/nwaples/rardecode"
	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
)

// VERSION 0.2.0
const VERSION = "0.2.1"

// RarExtractor ..
func RarExtractor(path, destination string) error {
	rr, err := rardecode.OpenReader(path, "")
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
			err = mkdir(filepath.Join(destination, header.Name))
			if err != nil {
				return err
			}
			continue
		}
		err = mkdir(filepath.Dir(filepath.Join(destination, header.Name)))
		if err != nil {
			return err
		}

		err = writeNewFile(filepath.Join(destination, header.Name), bufRr, header.Mode())
		if err != nil {
			return err
		}

	}

	return nil
}

// ZipExtractor ..
func ZipExtractor(source, destination string) error {
	r, err := zip.OpenReader(source)
	if err != nil {
		return err
	}
	defer r.Close()

	return unzipAll(&r.Reader, destination)
}

func unzipAll(r *zip.Reader, destination string) error {
	for _, zf := range r.File {
		if err := unzipFile(zf, destination); err != nil {
			return err
		}
	}

	return nil
}

func unzipFile(zf *zip.File, destination string) error {
	//nolint:gosec // Cleaned by filepath.Clean
	path := filepath.Join(destination, zf.Name)

	// Convert from ShiftJIS to UTF-8 (if needed)
	if strings.Contains(zf.Name, ".txt") {
		rc, err := zf.Open()
		if err != nil {
			return fmt.Errorf("%s: open compressed file: %v", zf.Name, err)
		}
		utf8Reader := transform.NewReader(rc, japanese.ShiftJIS.NewDecoder())
		bufRc := bufio.NewReader(utf8Reader)
		return writeNewFile(path, bufRc, zf.FileInfo().Mode())
	}

	// Clean the name for security's sake (avoid Zip Slip vulnerability)
	cleanPath := filepath.Clean(path)
	if !strings.HasPrefix(cleanPath, filepath.Clean(destination)) {
		return fmt.Errorf("%s: illegal file path", path)
	}

	if strings.HasSuffix(zf.Name, "/") {
		return mkdir(cleanPath)
	}

	rc, err := zf.Open()
	if err != nil {
		return fmt.Errorf("%s: open compressed file: %v", zf.Name, err)
	}
	defer rc.Close()

	bufRc := bufio.NewReader(rc)

	return writeNewFile(cleanPath, bufRc, zf.FileInfo().Mode())
}

func mkdir(dirPath string) error {
	err := os.MkdirAll(dirPath, 0o755)
	if err != nil {
		return fmt.Errorf("%s: making directory: %v", dirPath, err)
	}
	return nil
}

func writeNewFile(fpath string, in *bufio.Reader, fm os.FileMode) error {
	err := os.MkdirAll(filepath.Dir(fpath), 0o755)
	if err != nil {
		return fmt.Errorf("%s: making directory for file: %v", fpath, err)
	}

	out, err := os.Create(fpath)
	if err != nil {
		return fmt.Errorf("%s: creating new file: %v", fpath, err)
	}
	defer out.Close()

	err = out.Chmod(fm)
	if err != nil && runtime.GOOS != "windows" {
		return fmt.Errorf("%s: changing file mode: %v", fpath, err)
	}

	_, err = io.Copy(out, in)
	if err != nil {
		return fmt.Errorf("%s: writing file: %v", fpath, err)
	}
	return nil
}

func main() {
	i := flag.String("i", "", "path to rar file")
	o := flag.String("o", "", "destination path")
	flag.Parse()

	if *i == "" {
		fmt.Println("path is required")
		os.Exit(1)
	}

	if *o == "" {
		fmt.Println("destination is required")
		os.Exit(1)
	}

	err := RarExtractor(*i, *o)
	if err != nil {
		fmt.Errorf("failed to extract rar file: %v", err)
		os.Exit(1)
	}

	// display help if no arguments are passed
	if len(os.Args) == 1 {
		fmt.Println("Usage: gorar -i <path to file> -o <destination path>")
		os.Exit(1)
	}
}
