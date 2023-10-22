package main

import (
	"fmt"
	"runtime"
)

var m runtime.MemStats

type Task struct {
	todo string
}

func (task Task) DoSomething() {
	fmt.Println(task.todo)
}

func work() {
	workers := 5
	ch := make(chan *Task, workers)
	defer close(ch) // <- ** 最後にちゃんとcloseしないとリークする **

	for i := 0; i < workers; i++ {
		// 5つ並行タスクを処理する
		go func() {
			// rangeでタスクを待ち受ける。chがcloseされたら抜ける。
			for task := range ch {
				// タスクの処理を行う
				task.DoSomething()
			}
		}()
	}

	// タスクをチャンネルに送り込む処理
	for i := 0; i < 20; i++ {
		ch <- &Task{todo: fmt.Sprintf("todo_%d", i)}
	}
}

func PrintMemUsage() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)

	fmt.Printf("Alloc = %v MiB\t", bToMb(m.Alloc))
	fmt.Printf("TotalAlloc = %v MiB\t", bToMb(m.TotalAlloc))
	fmt.Printf("Sys = %v MiB\t", bToMb(m.Sys))
	fmt.Printf("Mallocs = %v MiB\t", bToMb(m.Mallocs))
	fmt.Printf("Frees = %v MiB\t", bToMb(m.Frees))
	fmt.Printf("HeapAlloc = %v MiB\t", bToMb(m.HeapAlloc))
	fmt.Printf("HeapSys = %v MiB\t", bToMb(m.HeapSys))
	fmt.Printf("NumGC = %v", m.NumGC)
	fmt.Printf("\n")
}

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}

func main() {
	// Goではruntimeパッケージを利用して各種メトリクスを取得できる。
	// 動作しているGoroutineの数を取得する「runtime.NumGoroutine」とメモリやGC(ガベージコレクション)の状況を取得する「runtime.ReadMemStats」は常に可視化して監視できるようにしたようが良い。

	// golang-stats-api-handlerというライブラリがあり、http経由でruntimeパッケージが提供される各種メトリクスを取得できる
	fmt.Printf("running goroutine is %d \n", runtime.NumGoroutine())
	PrintMemUsage()

	work()

	fmt.Printf("running goroutine is %d \n", runtime.NumGoroutine())
	PrintMemUsage()
}
