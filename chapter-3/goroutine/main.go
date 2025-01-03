package main

import (
	"sync"
	"fmt"
)

var wg sync.WaitGroup

func main() {
	queue := make(chan string)
	for i := 0; i < 2; i ++ { // 2つのgoroutine(ワーカー)を作成
		wg.Add(1)
		go fetchURL(queue)
	}

	queue <- "https://www.example.com"
	queue <- "https://www.example.net"
	queue <- "https://www.example.net/foo"
	queue <- "https://www.example.net/bar"

	close(queue) // goroutine に終了を伝える (※ チャンネルは一度しかcloseできない、close済みのチャンネルに対して送信を行うことはできない)
	wg.Wait() // 全てのgoroutineが終了するのを待つ
}

func fetchURL(queue chan string) {
	for {
		url, more := <- queue // closeされるとmoreがfalseになる
		if more {
			// url取得処理
			fmt.Println("fetching", url)
			// ...

		} else {
			fmt.Println("worker exit")
			wg.Done()
			return
		}
	}
}


func fetchURLWithRange(queue chan string, done chan<- bool) {
	for url := range queue {
		// url取得処理
		fmt.Println("fetching", url)
	}
	fmt.Println("worker exit")
	done <- true
}
