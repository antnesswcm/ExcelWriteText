package main

import (
	"fmt"
)

var letterToNum = map[rune]int{}
var fileList []string
var columnListNum []int
var columnList []string
var scopeList []string
var scopeListNum [][]int

var (
	files     string
	columns   string
	blankStop bool
	scopes    string
	extension string
	outpath   string
	help      bool
)

func main() {
	parseFlags() // 解析命令行参数
	println(blankStop)
	processParams() // 处理参数
	//processFiles()  // 处理文件

	fmt.Println("完成!")
}
