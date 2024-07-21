package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func which(fileName string) *string {
	path := os.Getenv("PATH")
	pathSplit := filepath.SplitList(path)
	for _, dir := range pathSplit {
		fullPath := filepath.Join(dir, fileName)
		fileInfo, err := os.Stat(fullPath)
		if err != nil {
			continue
		}

		mode := fileInfo.Mode()
		if !mode.IsRegular() {
			continue
		}

		if mode&0111 != 0 {
			return &fullPath
		}
	}
	return nil
}

func main() {
	var file string
	args := os.Args
	if len(args) < 2 {
		fmt.Println("The executable file should be defined")
	} else {
		file = args[1]
	}

	res := which(file)
	if res != nil {
		fmt.Println(*res)
	} else {
		fmt.Println("Not found")
	}
}
