package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func writeFile(filename, content string) error {
	if _, err := os.Stat(outpath); os.IsNotExist(err) {
		// 文件夹不存在，提示创建
		fmt.Printf("文件夹 %s 不存在，即将创建...\n", outpath)
		// 创建文件夹
		err := os.MkdirAll(outpath, os.ModePerm)
		if err != nil {
			fmt.Printf("创建文件夹 %s 失败：%s\n", outpath, err.Error())
			os.Exit(1)
		}
		fmt.Printf("文件夹 %s 创建成功！\n", outpath)
	}

	fullPath := filepath.Join(outpath, filename+"."+extension)

	// 创建文件
	file, err := os.Create(fullPath)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(content)
	if err != nil {
		return err
	}
	return nil
}
