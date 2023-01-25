package main

import (
	flag "github.com/spf13/pflag"

	"github.com/anonsachin/go-wizard/pkg/cli"
)

func main() {
	// Getting the inputs
	// defining the flags to be extracted
	moduleName := flag.String("module_name", "", "specifies the name of the golang module.")
	dirName := flag.String("dir_name", "", "specifies the name of the directory name to use while creating the project.")
	binaryDirName := flag.String("binary_dir_name", "", "specifies the name of the directory name for the binary directory to use while creating the project.")
	library := flag.Bool("library", false, "specifies if it is a library.")
	debug := flag.Bool("debug", false, "To run the cli in debug state.")
	flag.Parse()

	action := flag.Arg(0)
	module := flag.Arg(1)

	// the values of dirname can be inherited from the module name
	if *dirName == "" {
		if *moduleName == "" {
			if module != "" {
				*dirName = module
			}
		} else {
			*dirName = *moduleName
		}
	}

	if *binaryDirName == "" {
		if *moduleName == "" {
			if module != "" {
				*binaryDirName = module
			}
		} else {
			*binaryDirName = *moduleName
		}
	}

	args := &cli.ArgumentInformation{
		Args: flag.Args()[1:],
		StringFlags: map[string]string{
			"moduleName":    *moduleName,
			"dirName":       *dirName,
			"binaryDirName": *binaryDirName,
		},
		BoolFlags: map[string]bool{
			"library": *library,
			"debug": *debug,
		},
	}
	// specifying the handlers
	handler := &cli.Handler{
		Debug: args.BoolFlags["debug"],
	}
	// Strating execution
	handler.Router(action, args)
}
