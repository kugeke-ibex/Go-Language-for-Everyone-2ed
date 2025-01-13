package cat_test

import (
	"testing"
	cat "github.com/kugeke-ibex/go-language-for-everyone-2ed/chapter-6/Benchmark"
	"fmt"
)

// seedはベンチマーク用のトークンを作る
// 長さを受け取り、指定された長さの文字列のスライスを生成する
// 今回は、単純に"a"をn個並べたスライスを生成する
func seed(n int) []string {
	s := make([]string, n)
	for i := 0; i < n; i++ {
		s = append(s, "a")
	}
	return s
}

// benchはベンチマーク用のヘルパ
// テストしたい文字列の組み合わせ長と、文字列結合のための手続きを渡す
// それについてベンチマークを実行する
func bench(b *testing.B, n int, f func(...string) string) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		f(seed(n)...)
	}
}

func BenchmarkCat3(b *testing.B) {
	bench(b, 3, cat.Cat)
}

func BenchmarkBuf3(b *testing.B) {
	bench(b, 3, cat.Buf)
}

func BenchmarkCat100(b *testing.B) {
	bench(b, 100, cat.Cat)
}

func BenchmarkBuf100(b *testing.B) {
	bench(b, 100, cat.Buf)
}

func BenchmarkCat10000(b *testing.B) {
	bench(b, 10000, cat.Cat)
}

func BenchmarkBuf10000(b *testing.B) {
	bench(b, 10000, cat.Buf)
}

// ベンチマークのテストの実行
// go test -bench .
// 出力される内容は左から
// ループが実行された回数
// 1ループごとの所要時間
// 1ループごとのアロケーションされたバイト数
// 1ループごとのアロケーション回数を表示

// 実際にパフォーマンスチューニングをしていく場合には実装のどの部分で多くアロケーションされている確認したいので、その場合はpprofを使うのが便利
// Goプラグラムのパフォーマンスを測定する方法については「Profiling Go Programs - The Go Blog」を参照

// サブベンチマークの利用
func BenchmarkX(b *testing.B) {
	b.Run("n=3", func(b *testing.B) {
		bench(b, 3, cat.Cat)
		bench(b, 3, cat.Buf)
	})
	b.Run("n=100", func(b *testing.B) {
		bench(b, 100, cat.Cat)
		bench(b, 100, cat.Buf)
	})
	b.Run("n=10000", func(b *testing.B) {
		bench(b, 10000, cat.Cat)
		bench(b, 10000, cat.Buf)
	})
}

// 下記のようにサブベンチマークを利用することで、新しいベンチマークを追加するのが容易になる
// 様々なベンチマークを記述したい場合には、このようにサブベンチマークを利用すると便利
func BenchmarkConcatenate(b *testing.B) {
	benchCases := []struct {
		name string
		n int
		f func(...string) string
	}{
		{"Cat", 3, cat.Cat},
		{"Buf", 3, cat.Buf},
		{"Cat", 100, cat.Cat},
		{"Buf", 100, cat.Buf},
		{"Cat", 10000, cat.Cat},
		{"Buf", 10000, cat.Buf},
	}

	for _, bc := range benchCases {
		b.Run(fmt.Sprintf("%s-%d", bc.name, bc.n), func(b *testing.B) {
			bench(b, bc.n, bc.f)
		})
	}
}
