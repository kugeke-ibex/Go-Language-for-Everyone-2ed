package hex_text

import (
	"encoding/hex"
	"testing"
)

// encDecTestはエンコードとデコードのテストテーブルに関する型
type encDecTest struct {
	enc string
	dec []byte
}

// encDecTestsはテストケースのテーブルになっている
// Table Driven Tests の場合にはこのようにスライズによってテストケースを列挙
var encDecTests = []encDecTest{
	{"", []byte{}},
	{"0001020304050607", []byte{0, 1, 2, 3, 4, 5, 6, 7}},
	{"08090a0b0c0d0e0f", []byte{8, 9, 10, 11, 12, 13, 14, 15}},
	{"f0f1f2f3f4f5f6f7", []byte{0xf0, 0xf1, 0xf2, 0xf3, 0xf4, 0xf5, 0xf6, 0xf7}},
	{"f8f9fafbfcfdfeff", []byte{0xf8, 0xf9, 0xfa, 0xfb, 0xfc, 0xfd, 0xfe, 0xff}},
	{"67", []byte{'g'}},
	{"e3a1", []byte{0xe3, 0xa1}},
}

// TestEncodeはhexへのエンコードのテスト
func TestEncode(t *testing.T) {
	// 上で定義したテストケースのテーブルをイテレートしている
	for i, test := range encDecTests {
		dst := make([]byte, hex.EncodedLen(len(test.dec)))
		// エンコードした結果の長さを返す
		n := hex.Encode(dst, test.dec)
		if n != len(dst) {
			t.Errorf("#%d: bad return value: got: %d want: %d", i, n, len(dst))
		}
		// エンコードした結果をテストケースのエンコード結果と比較
		if string(dst) != test.enc {
			t.Errorf("#%d: got: %#v want: %#v", i, dst, test.enc)
		}
	}
}
