package cat

import (
	"bytes"
)

// Catは += 演算子を使って文字列を結合する
func Cat(ss ...string) string {
	var r string
	for _, s := range ss {
		r += s
	}

	return r
}

// Bufはbytes.Bufferを使って文字列を結合する
func Buf(ss ...string) string {
	var b bytes.Buffer
	for _, s := range ss {
		// NOTICE: エラーは無視している
		b.WriteString(s)
	}
	return b.String()
}
