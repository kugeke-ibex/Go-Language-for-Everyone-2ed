package main

import (
	"io"
	"net/http"
	"time"
)

func getHTTP(url string, dst io.Writer) error {
	client := &http.Client{
		// 10秒でタイムアウトする
		Timeout: 10 * time.Second,
	}

	req, _ := http.NewRequest("GET", url, nil)
	resp, err := client.Do(req)
	if err != nil {
		// レスポンスヘッダ取得までに10秒経過した場合にはここでエラー
		// request canceled (Client.Timeout exceeded while awaiting headers)

		return err
	}

	defer resp.Body.Close()

	_, err = io.Copy(dst, resp.Body)
	// ボディ取得完了までに10秒経過した場合はここでエラー
	// request canceled (Client.Timeout exceeded while awaiting headers)
	return err
}
