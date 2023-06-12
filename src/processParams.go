package main

import (
	"flag"
	"fmt"
	"os"
	"regexp"
	"strings"
	"unicode"
)

func processParams() {
	// 如果出现 -h 或 --help 标志，打印帮助信息并退出程序
	if help {
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
		for _, v := range columnList {
			// 判断字符串是否只包含一个字符
			if len(v) != 1 {
				fmt.Fprintf(os.Stderr, "%s参数解析失败！", v)
				os.Exit(1)
			}

			// 判断该字符是否为字母
			if !unicode.IsLetter(rune(v[0])) {
				fmt.Fprintf(os.Stderr, "%s参数解析失败！", v)
				os.Exit(1)
			}
			v = strings.ToLower(v)
			columnListNum = append(columnListNum, letterToNum[rune(v[0])])
		}
	} else {
		// 生成0 columnListNum
		for i := 0; i < len(fileList); i++ {
			columnListNum = append(columnListNum, 0)
		}
	}
	scopeList = strings.Split(scopes, ",")
	for _, v := range scopeList {
		r := regexp.MustCompile(`(?i)^[a-z]-[a-z]$`)
		switch {
		// 判断是否为空
		case v == "":
			scopeListNum = append(scopeListNum, []int{-1, -1})
		// 是否单个字母
		case len(v) == 1 && unicode.IsLetter(rune(v[0])):
			v = strings.ToLower(v)
			n := letterToNum[rune(v[0])]
			scopeListNum = append(scopeListNum, []int{n, -1})
		case r.MatchString(v):
			v = strings.ToLower(v)
			n1 := letterToNum[rune(v[0])]
			n2 := letterToNum[rune(v[2])]
			scopeListNum = append(scopeListNum, []int{n1, n2})

		default:
			fmt.Fprintf(os.Stderr, "%s参数解析失败！", v)
			os.Exit(1)
		}
	}
	if len(scopeListNum) > len(fileList) {
		fmt.Fprintf(os.Stderr, "指定的数据列数量大于处理的文件数量！")
		os.Exit(1)
	} else if len(scopeListNum) < len(fileList) {
		fmt.Fprintf(os.Stderr, "指定的数据列数量小于处理的文件数量！将自动匹配")
		for i := len(scopeListNum); i < len(fileList); i++ {
			scopeListNum = append(scopeListNum, []int{-1})
		}
	}
}
