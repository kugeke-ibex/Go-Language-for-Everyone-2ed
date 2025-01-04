package main

import (
	"fmt"
	"os"
)

func main() {
	f, _ := os.Open("file.txt")
	HandleData(f)
	HandleData(1)
	HandleData("hoge")
	HandleData(map[string]int{"foo": 1})
}

func HandleData(x interface{}) {
	// 型アサーション
	f, ok := x.(*os.File)
	if ok {
		// "file is file.txt"
		fmt.Println("file is " + f.Name())

		// 以下のようにos.Fileが満たすinterfaceへの型アサーションも可能
		// x.(io.Reader)
		// x.(io.Writer)
		// x.(io.Closer)
	}

	switch x.(type) {
	case int:
		// HandleData(1)のケース
		fmt.Printf("Number is %d\n", x)
	case string:
		// HandleData("hoge")
		fmt.Printf("Character is %s\n", x)
	case map[string]int:
		// HandleData(map[string]int{"foo": 1})
		fmt.Printf("Map is %v\n", x)
	// case map: // mapだけでは不完全な型なのでコンパイルエラー
	default:

	}
}
