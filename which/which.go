package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Please provide an argument")
		return
	}
	file := arguments[1]
	path := os.Getenv("PATH")
	pathSplit := filepath.SplitList(path)
	var results []string
	out:
	for _, directory := range pathSplit {
		fullPath := filepath.Join(directory, file)
		fileInfo, err := os.Stat(fullPath)
		if err != nil {
			continue
		}

		// Is it a regular file?
		mode := fileInfo.Mode()
		if !mode.IsRegular() {
			continue
		}

		// Is it executable?
		if mode&0111 != 0 {
			for _, result := range results {
				if result == fullPath {
					break out
				}
			}
			results = append(results, fullPath)
		}
	}
	for _, result := range results {
		fmt.Println(result)
	}
}

