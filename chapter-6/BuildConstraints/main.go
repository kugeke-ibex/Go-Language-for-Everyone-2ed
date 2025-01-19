//go:build hoge
// +build hoge

package main

import (
	"fmt"
)

// NOTE: Build Constraintsはpackageより前でかつ空行を開けて記述しなければならない

func main() {
	fmt.Println("Hello, World!")
}

// go vet によってbuildタグを検証。Go 1.10からはgo test実行時にも実行される
