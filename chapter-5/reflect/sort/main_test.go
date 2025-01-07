package main_test

import (
	"math/rand"
	"sort"
	"testing"

	. "github.com/kugeke-ibex/go-language-for-everyone-2ed/chapter-5/reflect/sort"
)

// ランダムな整数の配列を生成する
func randlist(n int) []int {
	// n個の配列を作り、そこに0からn-1を格納
	l := make([]int, n)
	for i := range l {
		l[i] = i
	}

	// 適当にシャッフル
	for i := range l {
		j := rand.Intn(i + 1)
		l[i], l[j] = l[j], l[i]
	}

	return l
}

// reflectを使ったソート
func BenchmarkSortReflect(b *testing.B) {
	master := randlist(25)
	l := make([]int, 25)

	for i := 0; i < b.N; i++ {
		copy(l, master)
		// マスターから作業用のバッファのlにデータをコピー

		// Sortwrapを使ってソートする
		sort.Sort(NewSortwrap(l, func(i, j int) bool {
			return l[i] < l[j]
		}))
	}
}

// reflectを使わす、組み込みの`sort.Ints`を使ったソート
func BenchmarkSortRaw(b *testing.B) {
	master := randlist(25)
	l := make([]int, 25)

	for i := 0; i < b.N; i++ {
		copy(l, master)
		// マスターから作業用のバッファのlにデータをコピー

		sort.Ints(l)
	}
}
