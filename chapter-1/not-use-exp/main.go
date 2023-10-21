package main

import (
	"fmt"
	"regexp"
	"strings"
)

// 正規表現パターンの生成にはコストがかかるので、初期化時に生成する。
var wordReg = regexp.MustCompile(`\w+`)

func main() {
	// 正規表現パッケージregexpパッケージは高機能だが、パフォーマンスに問題あり。
	// なので、文字列操作はstringsパッケージを基本使う。
	fmt.Println(strings.HasPrefix("hoge", "h"))
	fmt.Println(strings.HasSuffix("hoge", "e"))
	fmt.Println(strings.Contains("hoge", "og"))
	fmt.Println(strings.Fields("hoge piyo funnga PIYO"))
	fmt.Println(strings.Split("hoge piyo funnga PIYO", " "))
	fmt.Println(strings.SplitN("hoge,piyo,funnga,PIYO", ",", 4))
	fmt.Println(strings.TrimSpace(" pi yo "))
	fmt.Println(strings.Trim(",pi yo,", ","))
	fmt.Println(strings.Replace(",pi yo,", ",", "~", 1))

	// 実行中に使いたい場合は、MustXXXの関数を使わない。
	reg, err := regexp.Compile(`hello`)
	if err != nil {
		fmt.Println("regexp")
		fmt.Println(reg.Find([]byte("hello")))
	}
	fmt.Println("wordReg")
	fmt.Println(string(wordReg.Find([]byte("string hoge"))))
}
