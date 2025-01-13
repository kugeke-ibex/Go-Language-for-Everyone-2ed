package example_test

import (
	"fmt"
	"bufio"
	"os"
)


func ExampleHello() {
	// Pass
	fmt.Println("Hello")
	// Output: Hello

	// Fail
	// fmt.Println("World")
	// Output: world
}

func ExampleUnordered() {
	for _, v := range []int{1, 2, 3} {
		fmt.Println(v)
	}
	// Unordered outputを使うと、順不同でもテストが通る
	// Unordered output:
	// 2
	// 3
	// 1
}

// mapのイテレートは順不同なので、このテストはたまに失敗する
func ExampleShufullWillBeFailed() {
	x := map[string]int{"a": 1, "b": 2, "c": 3}
	for k, v := range x {
		fmt.Printf("k=%s: v=%d\n", k, v)
	}
	// Output:
	// k=a: v=1
	// k=b: v=2
	// k=c: v=3
}

// Unordered outputを使うと、順不同でもテストが通る
func ExampleShuffle() {
	x := map[string]int{"a": 1, "b": 2, "c": 3}
	for k, v := range x {
		fmt.Printf("k=%s: v=%d\n", k, v)
	}
	// Unordered output:
	// k=a: v=1
	// k=b: v=2
	// k=c: v=3
}

// src/bufio/example_test.goより

func ExampleWriter() {
	w := bufio.NewWriter(os.Stdout)
	fmt.Fprintf(w, "Hello")
	fmt.Fprintf(w, "World")
	w.Flush()  // Don't forget to flush
	// Output: Hello, World!
}

// Examplesはテストのための機能としてはシンプルなもの。
// ドキュメンテーションのために役に立つ機能であり、コードは編集可能で、実際に実行で実行できる例として価値がある。
// 自分で実装したライブラリのドキュメントに例を簡単に埋め込むことができ、かつそれらのコードが実行できるということがビルド時に保証される。
