package main_test

import (
	"testing"
	"reflect"
)

type T struct {
	x int
	ss []string
	m map[string]int
}

func TestStruct(t *testing.T) {
	ml := map[string]int{
		"a": 1,
		"b": 2,
	}

	t1 := T{
		x: 1,
		ss: []string{"a", "b"},
		m: ml,
	}

	t2 := T{
		x: 1,
		ss: []string{"a", "b"},
		m: ml,
	}

	if !reflect.DeepEqual(t1, t2) {
		t.Errorf("want %#v, got %#v", t1, t2)
	}
	// reflect.DeepEqualの振る舞い
	// ・型が異なればfalseを返す
	// ・arrayであればそれぞれのの値を再帰的に値を見て、比較していく
	// ・mapであれば、キーと値のペアを再帰的に比較していく
	// ・structであれば、フィールドの値を再帰的に比較していく
	// ・sliceであれば、長さとそれぞれの要素を再帰的に比較していく
	// ・pointerであれば、それぞれの指す値を再帰的に比較していく
	// ・interfaceであれば、実際の値が等しいか否かを比較する
	// ・funcであれば、それぞれの値を再帰的に比較していく
	// ・channelであれば、それぞれの値を再帰的に比較していく
	// ・complexであれば、それぞれの値を再帰的に比較していく
	// ・stringであれば、それぞれの値を再帰的に比較していく
}

type mapTest struct {
	a, b map[string]int
	eq bool
}

var mapTests = []mapTest{
	{map[string]int{"a": 1}, map[string]int{"b": 1}, false},
	{map[string]int{"a": 1}, map[string]int{"a": 1}, true},
}

func TestMapTable(t *testing.T) {
	for _, test := range mapTests {
		if r := reflect.DeepEqual(test.a, test.b); r != test.eq {
			t.Errorf("when a = %#v, and b = %#v, want %t, got %t", test.a, test.b, r, test.eq)
		}
	}
}
