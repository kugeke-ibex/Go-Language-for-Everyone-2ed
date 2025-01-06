package main

import (
	"os"
)

func main() {
	filename := "file.txt"
	f, err := os.Open(filename)
	if err != nil {
		return
	}

	defer f.Close()

	// ファイルから読み込んだバイト列をchに送る
	ch := make(chan []byte)
	go func() {
		defer close(ch)
		buf := make([]byte, 4096)
		for {
			n, err := f.Read(buf)
			if err != nil {
				return
			}
			ch <- buf[:n]
		}
	}()

	// 順次チャンネルからバイト列を読み込みそれを表示する
	for in := range ch {
		os.Stdout.Write(in)
	}
}
