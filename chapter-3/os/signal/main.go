package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	// "time"
)

func main () {
	defer fmt.Println("done")

	// 取り扱うシグナルを決める
	trapSignals := []os.Signal{
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT,
	}
	// 受信するためのチェンネルを用意
	sigCh := make(chan os.Signal, 1)
	// 受信する
	signal.Notify(sigCh, trapSignals...)

	// メインの処理を行う関数に渡す、キャンセル可能なコンテキストを作る
	ctx, cancel := context.WithCancel(context.Background())
	// 別goroutineでシグネルで待ち受ける
	go func() {
		// シグナルを受信するまでブロックする
		sig := <- sigCh
		fmt.Println("Got signal", sig)
		// シグナルで受信したので、終了させるためにキャンセル
		cancel()
	}()
	doMain(ctx)
}

func doMain(ctx context.Context) {
	defer fmt.Println("done doMain")
	for {
		select {
		case <- ctx.Done():
			return
		default:
		}
		// 何らかの処理
	}
}
