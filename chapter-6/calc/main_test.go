package main_test

import (
	"testing"
	calc "github.com/kugeke-ibex/go-language-for-everyone-2ed/chapter-6/calc"
)

// TestSumは加算のテストをする
// 引数には*testing.Tを渡す。
// 必ずTestから始まる名前にすると、go testでの実行対象になる。
// 個別にテストを実行する事も可能。 go test -run TestXXX
// Go 1.10以降だとテストがキャッシュされるので、明示的にクリアする場合は、go clean -testcache
func TestSum(t *testing.T) {
	// t.Fatalはテストが失敗したことを返すAPI
	// 多くのGoのテストコードでは条件分岐とt.Fatalを組み合わせて書くことになる。
	// t.Fatal以外にも、t.Fatalfもある(テスト失敗時のメッセージを加工する)
	if calc.Sum(1, 2) != 3 {
		t.Fatal("sum(1, 2) should be 3, but doesn't match")
	}
}
