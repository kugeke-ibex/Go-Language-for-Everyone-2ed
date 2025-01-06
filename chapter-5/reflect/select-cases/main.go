package main

import (
	"errors"
	"fmt"
	"os"
	"os/signal"
	"reflect"
	"syscall"
)

func main() {
	if err := _main(); err != nil {
		fmt.Printf("%s\n", err)
		os.Exit(1)
	}
}

func readFromFile(ch chan []byte, f *os.File) {
	defer close(ch) //全て終わったらチャンネルを閉じる
	defer f.Close() // 全て終わったらファイルを閉じる

	buf := make([]byte, 4096)
	for {
		// 読み込めるデータがあればそれをチャンネルに流す
		// ここではエラーがあってもファイルを追い続ける
		// (そうしないとio.EOFを受け取ったら、それ以上tailできなくなってしまう)
		if n, err := f.Read(buf); err == nil {
			ch <- buf[:n]
		}
	}
}

func makeChannelsForFiles(files ...string) ([]reflect.Value, error) {
	cs := make([]reflect.Value, len(files))

	for i, fn := range files {
		// データを流すようのチャンネルを作り・・・
		ch := make(chan []byte)

		// ファイルをオープン
		f, err := os.Open(fn)
		if err != nil {
			return nil, err
		}

		go readFromFile(ch, f)
		cs[i] = reflect.ValueOf(ch)
	}

	return cs, nil
}

func makeSelectCases(cs ...reflect.Value) ([]reflect.SelectCase, error) {
	// 与えられた分のchanの数だけreflect.SelectCaseを作成
	cases := make([]reflect.SelectCase, len(cs))
	for i, ch := range cs {
		// reflect.Valueの値がチャンネルでない場合はエラー
		if ch.Kind() != reflect.Chan {
			return nil, errors.New("argument must be channel")
		}

		// チャンネルの場合はSelectCaseを作成
		cases[i] = reflect.SelectCase {
			Chan: ch,
			Dir: reflect.SelectRecv,
		}
	}

	return cases, nil
}

func doSelect(cases []reflect.SelectCase) {
	for {
		// 基本はreflect.SelectCaseを複数作り、
		// reflect.Selectにそれを渡すことによって動的にselectと同様に
		// チャンネルへのデータを持つことができるようになります。
		/**
			reflect.Select([]reflect.SelectCase{
				reflect.SelectCase{ ... }, // 0番目のcase
				reflect.SelectCase{ ... }, // 1番目のcase
				reflect.SelectCase{ ... }, // 2番目のcase
			})
		*/
		if chosen, recv, ok := reflect.Select(cases); ok {
			fmt.Printf("\n=== %s ===\n%s", os.Args[chosen+1], recv.Interface())
		}
	}
}

func _main() error {
	if len(os.Args) < 2 {
		return errors.New("prog [file1 file2 ...]")
	}

	// シグナル処理のおまじない
	sigch := make(chan os.Signal, 1)
	signal.Notify(sigch, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)

	// os.Argsの最初の引数はこのコマンド名
	// それを除外した分が対象ファイル名
	channels, err := makeChannelsForFiles(os.Args[1:]...)
	if err != nil {
		return err
	}

	// 上記ループで得たチャンネルから動的に
	// select caseを作成する
	cases, err := makeSelectCases(channels...)
	if err != nil {
		return err
	}

	// selectを動的に作成・実行する
	go doSelect(cases)

	// シグナルを受け取るまでブロックし続ける
	// Ctrl-Cとうつと終了
	select {
	case <-sigch:
		return nil
	}

	return nil
}
