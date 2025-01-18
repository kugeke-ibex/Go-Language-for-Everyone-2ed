package main

import "fmt"
/**
	競合状態を検出するためには、以下のコマンドを実行する
	go run -race main.go
*/
func main() {
	c := make(chan bool)
	// 排他制御内に並行にmapにアクセスすると競合状態が発生する
	// goroutineからの書き込みとmain goroutineからの書き込みを並行に行って、競合状態を発生させる
	m := make(map[string]string)
	go func() {
		m["1"] = "a" // 1つ目の競合するメモリアクセス
		c <- true
	}()
	m["2"] = "b" // 2つ目の競合するメモリアクセス
	// cはunbuffered channelなのでchannelへのsendが完了されるまで待つ
	/**
		1. Unbuffered Channel
		特徴
			•	バッファを持たないチャンネル。
			•	送信（send）と受信（receive）は、同時に行われる必要がある。
			•	送信側（ch <- value）は、受信側（<- ch）が値を受け取るまでブロック（待機）する。
			•	受信側も、送信側が値を送るまでブロックする。

		2. Buffered Channel
		特徴
			•	バッファを持つチャンネル。
			•	バッファサイズを指定して作成する（例: make(chan int, 3)）。
			•	する。
			•	送信側（ch <- value）は、チャンネルのバッファに空きがある場合はブロックせずに送信できる。
			•	受信側（<- ch）は、チャンネルにデータが存在する場合はブロックせずに受信できる。
	*/

	<-c
	for k, v := range m {
		fmt.Println(k, v)
	}
}
