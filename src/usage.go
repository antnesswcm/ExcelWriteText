package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

func usage() {
	fmt.Fprintf(os.Stderr, "用法: %s [选项] [files]\n", filepath.Base(os.Args[0]))
	fmt.Fprintln(os.Stderr, "选项:")
	flag.PrintDefaults()
}
