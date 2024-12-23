package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/mattn/go-isatty"
)

// Flush() error を実装しているインターフェイスを定義
type flusher interface {
	Flush() error
}

func main() {
	var output io.Writer

	if isatty.IsTerminal(os.Stdout.Fd()){
		fmt.Print("hoge")
		// 標準入力が端末なら出力先はos.Stdoutそのも
		output = os.Stdout
		} else {
		fmt.Print("piyo")
		// 標準入力が端末でなければbufio.Writerでラップ
		output = bufio.NewWriter(os.Stdout)
	}

	for i := 0; i < 100; i++ {
		fmt.Fprintln(output, strings.Repeat("x", 100))
	}

	if _o, ok := output.(flusher); ok {
		// Flush()を実装している場合 (bufio.Writer)のみFlush()を行う。
		_o.Flush()
	}
}
