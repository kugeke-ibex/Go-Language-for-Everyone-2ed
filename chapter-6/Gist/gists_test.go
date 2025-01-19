package main_test

import (
	"io"
	"strings"
	"testing"
	. "github.com/kugeke-ibex/go-language-for-everyone-2ed/chapter-6/Gist" // メインパッケージをインポート
)

func TestGetGists(t *testing.T) {
	// ここでDoGistRequestを上書き
	// このDoGistRequestはTestGetGistsのスコープの中でのみ有効
	// 上書きすることにより、ListGistsの挙動をかえ、ダミーの実装を実現できる
	DoGistRequest = func(user string) (io.Reader, error) {
		return strings.NewReader(`
			[
				{"html_url": "https://gist.github.com/mattn/1234567890"},
				{"html_url": "https://gist.github.com/mattn/0987654321"}
			]
		`), nil
	}
	urls, err := ListGists("mattn")
	if err != nil {
		t.Fatalf("list gists caused error: %s", err)
	}
	if expected := 2; len(urls) != expected {
		t.Fatalf("want %d, got %d", expected, len(urls))
	}
}
