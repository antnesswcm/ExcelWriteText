package main

import (
	"fmt"
	"os"
)

func processFiles() {
	for i, file := range fileList {
		column := columnListNum[i] // 同时遍历columnListNum
		err := processFile(file, column)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			continue
		}
	}
}
