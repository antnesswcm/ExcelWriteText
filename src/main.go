package main

import (
	"fmt"
)

var letterToNum = map[rune]int{}
var fileList []string
var columnListNum []int
var columnList []string

var (
	files     string
	columns   string
	blankStop bool
	scope     string
	extension string
	outpath   string
	help      bool
	helpLong  bool
)

func main() {
	parseFlags()    // 解析命令行参数
	processParams() // 处理参数
	processFiles()  // 处理文件

	fmt.Println("完成!")
}
