package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

func parseFlags() {
	flag.StringVar(&files, "f", "", "以逗号分隔的文件列表，不可与[files]同时使用")

	flag.StringVar(&columns, "c", "", "以逗号分隔的列地址列表如:\"-c a,b\"")
	flag.StringVar(&scopes, "s", "", "以逗号分隔的列地址范围如:\"-s a-b,c-d\"")

	flag.StringVar(&extension, "e", "txt", "指定文件后缀名")
	flag.StringVar(&outpath, "o", "./output", "指定输出文件夹")

	flag.BoolVar(&blankStop, "b", false, "遇见空单元停止")
	flag.BoolVar(&help, "h", false, "显示帮助信息")
	flag.BoolVar(&help, "help", false, "显示帮助信息")

	// 自定义打印帮助信息的格式
	flag.Usage = usage

	flag.Parse()
	// 游离参数
	fileList = flag.Args()
}
func usage() {
	fmt.Fprintf(os.Stderr, "用法: %s [选项] [files]\n", filepath.Base(os.Args[0]))
	fmt.Fprintln(os.Stderr, "选项:")
	flag.PrintDefaults()
}
