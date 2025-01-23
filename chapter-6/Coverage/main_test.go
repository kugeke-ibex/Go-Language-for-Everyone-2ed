package main_test

import (
	"testing"
	cov "github.com/kugeke-ibex/go-language-for-everyone-2ed/chapter-6/Coverage"
)

type Case struct {
	in, out string
}

var cases = []Case{
	{"", "wordless?"},
	{"今日は天気ですね", "one word"},
	{"Go Go Go Go Go Go Go!", "many words"},
}

func TestWordCount(t *testing.T) {
	for i, c := range cases {
		w := cov.Words(c.in)
		if w != c.out {
			t.Errorf("#%d: Words(%s) got %s; want %s", i, c.in, w, c.out)
		}
	}
}

// カバレッジを取得するには、 go test -cover を実行する
// より詳しい結果のカバレッジ表示するには、 go test -coverprofile=coverage.out を実行する
// go tool cover -func=coverage.out を実行すると、カバレッジの詳細が表示される
// go tool cover -html=coverage.out を実行すると、カバレッジの詳細がHTMLで表示される
// go test -covermode=count -coverprofile=count.out を実行すると、カバレッジの詳細がカウントで表示される
