package cli

import (
	"fmt"
	"os"
	"path/filepath"
)

var (
	main = `package main

func main(){

}
`
	ignore = `
/vendor
`
)

func (h *Handler) SetupFiles(dirName string, binaryDirName string, lib bool) error {
	// making the main directory
	err := os.MkdirAll(filepath.Join(dirName), os.ModePerm)

	if err != nil {
		if h.Debug {
			fmt.Printf("Unable to make base directory %s \n",dirName)
		}
		return err
	}

	// making the command directory
	if !lib {
		err = os.MkdirAll(fmt.Sprintf(filepath.Join(dirName, "cmd", binaryDirName)), os.ModePerm)

		if err != nil {
			if h.Debug {
				fmt.Println("Unable to make cmd directory")
			}
			return err
		}
	}

	// making the pkg directory
	err = os.MkdirAll(fmt.Sprintf(filepath.Join(dirName, "pkg")), os.ModePerm)

	if err != nil {
		if h.Debug {
			fmt.Println("Unable to make pkg directory")
		}
		return err
	}

	// main file with empty func
	mainFile, err := os.Create(filepath.Join(dirName, "cmd", binaryDirName, "main.go"))

	if err != nil {
		if h.Debug {
			fmt.Println("Unable to make main file")
		}
		return err
	}

	// writing contents
	_, err = mainFile.Write([]byte(main))

	// gitignore file
	ignoreFile, err := os.Create(filepath.Join(dirName, ".gitignore"))

	if err != nil {
		if h.Debug {
			fmt.Println("Unable to make .gitignore file")
		}
		return err
	}

	// writing contents
	_, err = ignoreFile.Write([]byte(ignore))

	return nil
}
