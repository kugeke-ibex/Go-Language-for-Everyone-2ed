package main

import (
	"flag"
	"fmt"
)

var version string = "1.0.0"

func main() {
	var showVersion bool
	// -v -versionが指定された場合にshowVersion変数が真になるように定義
	flag.BoolVar(&showVersion, "v", false, "show version")
	flag.BoolVar(&showVersion, "version", false, "show version")
	flag.Parse() // 引数からオプションをパースする
	if showVersion {
		// バージョン番号を表示して終了
		fmt.Println("version:", version)
		return
	}
	// ...
}
