package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func which(fileName string) []string {
	var result []string
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
			result = append(result, fullPath)
		}
	}
	return result
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
	if len(res) > 0 {
		for _, path := range res {
			fmt.Println(path)
		}
	} else {
		fmt.Println("Not found")
	}
}
