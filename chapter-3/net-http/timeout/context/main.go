package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"time"
)

func getHTTP(url string, dst io.Writer) error {
	// 10秒でタイムアウトする Context を作る
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	// context を与えたリクエストを使って実行
	resp, err := client.Do(req.WithContext(ctx))
	if err != nil {
		// レスポンスヘッダ取得までに10秒経過した場合にはここでエラー
		// context deadline exceeded
		return err
	}
	defer resp.Body.Close()
	_, err = io.Copy(dst, resp.Body)
	// ボディ取得完了までに10秒経過した場合はここでエラー
	// context deadline exceeded
	return err

	// キャンセル可能なコンテキストを作る
	ctx, cancel = context.WithCancel(context.Background())
	defer cancel()
	// ctxを親にした、1秒でタイムアウトするコンテキストを作る
	ctxChild, cancelChild := context.WithTimeout(ctx, time.Second)
	defer cancelChild()

	// どちらかのコンテキストが完了するまで待つ
	select {
		case <- ctxChild.Done():
			fmt.Println("child", ctxChild.Err())
		case <- ctx.Done():
			fmt.Println("parent", ctx.Err())
	}
	return nil
}

type Context interface {
	// デッドライン時刻、デッドラインが設定されているか
	Deadline() (deadline time.Time, ok bool)

	// 完了 (タイムアウト、キャンセルなどを含む)を知らせるチャンネル
	Done() <-chan struct{}

	// 完了した場合の完了理由を保持しているエラー値
	Err() error

	// 保持している任意型の値を返す
	Value(key interface{}) interface{}
}
