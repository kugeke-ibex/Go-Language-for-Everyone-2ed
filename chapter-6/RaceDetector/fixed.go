package main

import (
	"fmt"
	"sync"
)

/**
競合状態を検出するためには、以下のコマンドを実行する
go run -race fixed.go
*/
func fixed() {
	var mux sync.RWMutex
	c := make(chan bool)
	m := make(map[string]string)
	go func() {
		mux.Lock()
		m["1"] = "a"
		mux.Unlock()
		c <- true
	}()
	mux.Lock()
	m["2"] = "b"
	mux.Unlock()


	<-c
	for k, v := range m {
		fmt.Println(k, v)
	}
}

func main() {
	fixed()
}
