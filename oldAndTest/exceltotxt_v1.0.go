package main

import (
	"fmt"
	"os"

	"github.com/tealeg/xlsx"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("用法: ExcelToTxt <filename.xlsx>")
		os.Exit(1)
	}
	for _, v := range os.Args[1:] {
		work(v)
	}
	fmt.Println("完成!")
}
func work(filename string) {
	xlFile, err := xlsx.OpenFile(filename)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for _, sheet := range xlFile.Sheets {
		for _, row := range sheet.Rows {
			// 读取第一列作为文件名，第二列作为文件内容
			fileName := row.Cells[0].Value
			fileContent := row.Cells[1].Value

			// 将文件内容写入文件
			f, err := os.Create(fileName + ".txt")
			if err != nil {
				fmt.Println(err)
				continue
			}
			_, err = f.WriteString(fileContent)
			if err != nil {
				fmt.Println(err)
			}
			f.Close()
		}
	}
}
