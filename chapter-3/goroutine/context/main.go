package main

import (
	"fmt"
	"sync"
	"context"
)

var wg sync.WaitGroup

func main() {
	// キャンセルするためのContextを生成
	ctx, cancel := context.WithCancel(context.Background())
	queue := make(chan string)
	for i := 0; i < 2; i ++ {
		wg.Add(1)
		go fetchURL(ctx, queue)
	}

	queue <- "https://www.example.com"
	queue <- "https://www.example.net"
	queue <- "https://www.example.net/foo"
	queue <- "https://www.example.net/bar"

	cancel() // ctxを終了させる
	wg.Wait() // 全てのgoroutineが終了するのを待つ
}

// 処理の中断やタイムアウトにはcontextパッケージを利用することが一般的なため、自作の処理でも途中で打ち切りやタイムアウトをサポートしたい場合は、第一引数にcontext.Contextを受け取るインターフェイスとすべき。
func fetchURL(ctx context.Context, queue chan string) {
	for {
		select {
		case <- ctx.Done():
			fmt.Println("worker exit")
			wg.Done()
			return
		case url := <- queue:
			// URL取得処理
		}
	}
}
