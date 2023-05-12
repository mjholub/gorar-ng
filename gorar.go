package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/samber/lo"
)

func main() {
	i := flag.FlagSet("i", "", "path to rar file")
	o := flag.FlagSet("o", "", "destination path")
	flag.Parse()

	switch len(os.Args) {
	// fall back to treating one argument as the input file
	// and the latter as the destination unless they have *zip or *rar extensions
	case 0:
		fmt.Println("Usage: gorar -i <path to file> -o <destination path>")
		os.Exit(1)
	case 1, *i == "":
		if strings.HasSuffix(*i, ".zip") || strings.HasSuffix(*i, ".rar") {
			*i = os.Args[1]
		}
		fmt.Println("No input file specified")
		os.Exit(1)
	case 1, *o == "":
		o = strings.TrimSuffix(*i, filepath.Ext(*i))
	default:
		for arg := range os.Args {
			inputArchives := []string{".zip", ".rar"}
			if strings.HasSuffix(arg, inputArchives...)
			 {
				inputArchives = append(inputArchives, arg)
				*i = lo.ForEach 
				if *o == "" {
					*o = strings.TrimSuffix(*i, filepath.Ext(*i))
				}
			}
		}
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
