package test_main

import (
	"testing"
	"os"
)

func TestMain(m *testing.M) {
	setup() //何らかの初期化処理
	exitCode := m.Run() // 対象のテストケースを実行し、実行結果をexit codeとして返却
	showdown() //何らかの終了処理
	// exit codeをos.Exitに渡すことでテスト全体の終了コードを通知
	os.Exit(exitCode)

}
