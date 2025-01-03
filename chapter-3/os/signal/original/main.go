package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// os.Signalインターフェイス
// type Signal interface {
// 	String string
// 	Signal() // to distinguish from other Stringers
// }

type MySignal struct {
	message string
}

func (s MySignal) String() string {
	return s.message
}

func (s MySignal) Signal() {}

func main() {
	log.Println("[info] Start")
	trapSignals := []os.Signal{
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT,
	}
	// 受信するためのチェンネルを用意
	sigCh := make(chan os.Signal, 1)

	// 10秒後にstgChにMySignalの値を送信
	time.AfterFunc(10*time.Second, func() {
		sigCh <- MySignal{"time out"}
	})


	signal.Notify(sigCh, trapSignals...)
	// 受信するまで待ち続ける
	sig := <- sigCh
	switch s := sig.(type) { // 型アサーションで判別
	case syscall.Signal:
		// osからのシグナルの場合
		log.Printf("[info] Got signal: %s(%d)", s, s)
	case MySignal:
		// アプリケーション独自のシグナルの場合
		log.Printf("[info] %s", s) //.String()が評価される
	}
}


