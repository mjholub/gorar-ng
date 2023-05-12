package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	iohandlers "github.com/154pinkchairs/gorar-ng/handlers/io"
	"github.com/154pinkchairs/gorar-ng/handlers/rar"
	"github.com/154pinkchairs/gorar-ng/handlers/zip"
)

type ArchiveHandler interface {
	Extract(*iohandlers.Dir) error
}

type inputFiles []string

func (i *inputFiles) String() string {
	return fmt.Sprint(*i)
}

func (i *inputFiles) Set(value string) error {
	*i = append(*i, value)
	return nil
}

func main() {
	var inputs inputFiles
	flag.Var(&inputs, "i", "paths to rar files")
	o := flag.String("o", "", "destination path")
	flag.Parse()

	if len(inputs) == 0 && len(os.Args) > 1 {
		inputs = append(inputs, os.Args[1])
	}

	if *o == "" && len(os.Args) > 2 {
		*o = os.Args[2]
	}

	if len(inputs) == 0 {
		fmt.Println(
			"No input files specified. Usage: gorar -i <path to file1> -i <path to file2> -o <destination path>")
		os.Exit(1)
	}

	if *o == "" {
		fmt.Println("Destination is required. Usage: gorar -i <path to file1> -i <path to file2> -o <destination path>")
		os.Exit(1)
	}

	for _, i := range inputs {
		dir := iohandlers.Dir{
			Path: *o,
		}
		var handler ArchiveHandler
		switch {
		case strings.HasSuffix(i, ".rar"):
			handler = &rar.Archive{Path: i}
		case strings.HasSuffix(i, ".zip"):
			handler = &ziphandlers.ZipArchive{Path: i}
		default:
			fmt.Printf("Unsupported file type for file %s\n", i)
			os.Exit(1)
		}

		err := handler.Extract(&dir)
		if err != nil {
			fmt.Printf("Failed to extract archive file %s: %v\n", i, err)
			os.Exit(1)
		}
	}
}
