package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"unicode"
)

var letterToNum = map[rune]int{}
var columnListNum []int
var extension string
var outpath string

func main() {
	var file string
	var columns string
	var columnList []string

	flag.StringVar(&file, "f", "", "以逗号分隔的文件列表，不可与[files]同时使用")
	flag.StringVar(&columns, "c", "", "以逗号分隔的列地址列表如:\"-c a\"")
	flag.StringVar(&extension, "e", "txt", "指定文件后缀名")
	flag.StringVar(&outpath, "o", "", "指定输出文件夹")

	// 注册 -h 和 --help 标志参数
	help := flag.Bool("h", false, "显示帮助信息")
	helpLong := flag.Bool("help", false, "显示帮助信息")

	// 自定义打印帮助信息的格式
	flag.Usage = usage

	flag.Parse()

	// 如果出现 -h 或 --help 标志，打印帮助信息并退出程序
	if *help || *helpLong {
		flag.Usage()
		os.Exit(0)
	}
	if len(os.Args) < 2 {
		flag.Usage()
		fmt.Println("=================")
		fmt.Println("交互式还未完成！")
		fmt.Scanln()
		os.Exit(1)
	}
	files := flag.Args()
	if outpath == "" {
		path, err := os.Getwd()
		if err != nil {
			fmt.Println("没有指定输出文件夹且获取当前文件夹失败：\n", err)
			os.Exit(1)
		}
		outpath = path
	}

	//fmt.Println("游离: ", files)
	//fmt.Println("-f: " + file)
	//fmt.Println("-c: " + columns)

	//file flag处理
	if (file != "") == (len(files) != 0) {
		fmt.Fprintf(os.Stderr, "请不要同时给定位置参数与-F标志参数！")
		os.Exit(1)
	} else if file != "" {
		files = strings.Split(file, ",")
	}
	//fmt.Println("files: ", files)
	for i, ch := 'a', 0; i <= 'z'; i, ch = i+1, ch+1 {
		letterToNum[i] = ch
	}
	// columns flag处理
	if columns != "" {
		columnList = strings.Split(columns, ",")
		if len(columnList) != len(files) {
			fmt.Fprintf(os.Stderr, "指定的列位置与文件数量不匹配！")
			os.Exit(1)
		}
		// 使用映射关系将字符串映射为数字
		for _, i := range columnList {
			// 判断字符串是否只包含一个字符
			if len(i) != 1 {
				fmt.Fprintf(os.Stderr, "%s参数解析失败！", i)
				os.Exit(1)
			}

			// 判断该字符是否为字母
			if !unicode.IsLetter(rune(i[0])) {
				fmt.Fprintf(os.Stderr, "%s参数解析失败！", i)
				os.Exit(1)
			}
			i = strings.ToLower(i)
			columnListNum = append(columnListNum, letterToNum[rune(i[0])])
		}
	} else {
		// 生成0 columnListNum
		for i := 0; i < len(files); i++ {
			columnListNum = append(columnListNum, 0)
		}
	}
	//fmt.Println("columnList: ", columnList)

	//fmt.Println("columnListNum", columnListNum)

	for i, file := range files {
		column := columnListNum[i]
		err := processFile(file, column)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	}
	fmt.Println("完成!")
}
