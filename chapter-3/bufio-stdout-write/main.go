package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	b := bufio.NewWriter(os.Stdout)
	// 標準出力をラップするbufio.Writerを作成
	for i := 0; i < 100; i++ {
		// bufio.Writeに対して書き込みを行う      f
		fmt.Fprintln(b, strings.Repeat("x", 100))
	}

	b.Flush()

	// 64KBのバッファを持つWriterを作る
	// b1 := bufio.NewWriterSize(w, 65536)
	// 4096byteのバッファを持つWriterを作ろうとするとb1がそのまま返る
	// b2 := bufio.NewWriterSize(b1, 4096)
}
