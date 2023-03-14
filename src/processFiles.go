package main

import (
	"fmt"
	"github.com/tealeg/xlsx"
	"os"
)

func processFiles() {
	for i, file := range fileList {
		// todo 使用并发处理每个文件
		column := columnListNum[i] // 同时遍历columnListNum
		scope := scopeListNum[i]   // 同时遍历scopeListNum

		err := processFile(file, column, scope)
		if err != nil {
			fmt.Fprintf(os.Stderr, "文件%s读取失败:\n%s\n", file, err)
			continue
		}
	}
}

func processFile(filename string, column int, scope []int) (err error) {
	var fileName string
	var content string
	xlFile, err := xlsx.OpenFile(filename)
	if err != nil {
		return err
	}
	for _, sheet := range xlFile.Sheets {
		//rows := len(sheet.Rows)
		emptyRow := 0
		for _, row := range sheet.Rows {
			// 判断是否为空行
			if len(row.Cells) == 0 {
				emptyRow++
				continue // 跳过空行
			}
			//println(len(row.Cells))
			//continue
			// scope 参数处理
			if scope[0] == -1 {
				scope[0] = column + 1
			}
			if scope[1] == -1 {
				scope[1] = len(row.Cells) - 1
			}
			// 判断column与scope是否出界
			if column > len(row.Cells)-1 {
				// todo 说明
				break
			}
			if scope[0] > scope[1] {
				// todo 说明
				break
			} else if len(row.Cells)-1 < scope[1] {
				// todo 说明
				break
			}

			// 读取文件名
			fileName = row.Cells[column].Value
			if fileName == "" {
				continue
			}

			// 读取内容
			content = ""
			for i := scope[0]; i <= scope[1]; i++ {
				fileContent := row.Cells[i].Value
				if fileContent != "" {
					content += fileContent
				} else if blankStop {
					break
				}
			}
			// 处理写入
			if content != "" {
				err = writeFile(fileName, content)
				if err != nil {
					fmt.Fprintln(os.Stderr, err)
					continue
				}
			}
		}
	}
	return nil
}
