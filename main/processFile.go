package main

import (
	"fmt"
	"github.com/tealeg/xlsx"
	"os"
)

func processFile(filename string, column int) (err error) {
	var fileName string
	var content string
	xlFile, err := xlsx.OpenFile(filename)
	if err != nil {
		return err
	}

	for _, sheet := range xlFile.Sheets {
		for _, row := range sheet.Rows {
			// 读取第column作为文件名，后面的列作为文件内容
			fileNameCell := row.Cells[column]
			if fileNameCell.Value != "" {
				fileName = fileNameCell.Value
			} else {
				continue
			}
			content = ""
			for i := column + 1; i < len(row.Cells); i++ {
				fileContentCell := row.Cells[i]
				if fileContentCell.Value != "" {
					fileContent := fileContentCell.Value
					content += fileContent
				} else {
					break
				}
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
