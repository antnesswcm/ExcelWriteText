package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func processParams() {
	// 如果出现 -h 或 --help 标志，打印帮助信息并退出程序
	if help || helpLong {
		flag.Usage()
		os.Exit(0)
	}

	// 交互式
	if len(os.Args) < 2 {
		flag.Usage()
		fmt.Println("=================")
		fmt.Println("交互式还未完成！")
		fmt.Scanln()
		os.Exit(1)
	}

	// -o 参数处理
	if outpath == "" {
		path, err := os.Getwd()
		if err != nil {
			fmt.Println("没有指定输出文件夹且获取当前文件夹失败：\n", err)
			os.Exit(1)
		}
		outpath = path
	}

	//-f 参数处理
	if (files != "") == (len(fileList) != 0) {
		fmt.Fprintf(os.Stderr, "请不要同时给定files参数与-f标志参数！")
		os.Exit(1)
	} else if files != "" {
		fileList = strings.Split(files, ",")
	}

	// 创建字母映射数字表
	for i, ch := 'a', 0; i <= 'z'; i, ch = i+1, ch+1 {
		letterToNum[i] = ch
	}

	// columns 参数处理
	if columns != "" {
		columnList = strings.Split(columns, ",")
		if len(columnList) != len(fileList) {
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
		for i := 0; i < len(fileList); i++ {
			columnListNum = append(columnListNum, 0)
		}
	}
}
