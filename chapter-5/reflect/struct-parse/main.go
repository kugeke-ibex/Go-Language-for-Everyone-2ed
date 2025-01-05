package main

import (
	"fmt"
	"reflect"
	"strings"
)

// struct tagでJSONオブジェクトに変換された際のキー名を指定できる
type Point struct {
	X int `json:"x,omitempty"`
	Y int `json:"y,omitempty"`
}

func main() {
	t := reflect.TypeOf(Point{X: 10, Y:20})
	for i := 0; i < t.NumField(); i ++ {
		f := t.Field(i)
		if f.PkgPath != "" {
			// If PkgPath is non empty
			// then it's an unexported field
			continue
		}

		tag := f.Tag.Get("json")
		// tagはstringなので、これに対して好きな操作を行う
		fmt.Println(tag)

		// タグの処理イメージ
		if v, ok := f.Tag.Lookup("json"); ok {
			parts := strings.Split(v, ",")
			// url encodeした際のフィールド名
			// 値がゼロ値だった場合の挙動
			// ヒント
			name := parts[0]
			omitempty := parts[1] // "omitempty" か ""
			// hint := parts[2]
			fmt.Println(name, omitempty)
		}
	}
}
