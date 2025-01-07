package performance_test

import (
	"reflect"
	"testing"
)

// アクセス速度を比較するためのstruct
type StructAccess struct {
	Int int
}

// こちらの関数では「もしinterface{}な値を与えられたときに
// 特定の型だったらその要素を得る」という処理をreflectを使って行った場合のコードを計測
func BenchmarkDetectTypeReflect(b *testing.B) {
	var s interface{} = StructAccess{Int: 100}
	for i := 0; i < b.N; i++ {
		rv := reflect.ValueOf(s)
		if rv.Type().Name() == "StructAccess" {
			_ = s.(StructAccess).Int
		}
	}
}

// こちらは最初から(コンパイル時から)型がわかってさえいれば そのままアクセスできるので、その場合の通常Goコードを計測
func BenchmarkDetectNone(b *testing.B) {
	s := StructAccess{Int: 100}
	for i := 0; i < b.N; i++ {
		_ = s.Int
	}
}


// このコードではreflectを使わず、型アサーションのみで処理している
func BenchmarkDetectTypeAssert(b *testing.B) {
	var s interface {} = StructAccess{Int: 100}
	for i := 0; i < b.N; i++ {
		if sa, ok := s.(StructAccess); ok {
			_ = sa.Int
		}
	}
}
// ベンチマークに関するテストを実行
// go test -bench .  -benchmem
