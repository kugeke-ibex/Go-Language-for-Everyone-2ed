package main

import (
	"fmt"
	"reflect"
	"sort"
)

type IntSlice []int

// 上記の型について、Len、Less、Swapを実装することで、sort.Interfaceというインターフェースを満たす。
func (p IntSlice) Len() int {
	return len(p)
}

func (p IntSlice) Less(i, j int) bool {
	return p[i] < p[j]
}

func (p IntSlice) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

type Sortwrap struct {
	value reflect.Value
	// 任意の配列が入ったreflect.Value
	lessfunc func(int, int) bool
	// 上記value内の値の比較を行うための関数
}

// コンストラクタ。配列と、その要素を比較する関数を受け取る
func NewSortwrap(v interface{}, lessfunc func(int, int) bool) *Sortwrap {
	// ここではvが本当に配列かどうかの確認を必要となるが、省略
	return &Sortwrap{
		value: reflect.ValueOf(v),
		lessfunc: lessfunc,
	}
}

// sort.InterfaceのLen()を満たすメソッド
// 単純に格納している配列の長さを返す
func (s *Sortwrap) Len() int {
	return s.value.Len()
}

// sort.InterfaceのLess()を満たすメソッド
// コンストラクタで与えられた関数を実行している
func (s *Sortwrap) Less(i, j int) bool {
	return s.lessfunc(i, j)
}

// sort.InterfaceのSwap()を満たすメソッド
// i番目とj番目の要素を入れ替える
// この操作は全てreflectで行なっている
func (s *Sortwrap) Swap(i, j int) {
	value := s.value
	v1 := value.Index(i).Interface()
	v2 := value.Index(j).Interface()
	value.Index(j).Set(reflect.ValueOf(v1))
	value.Index(i).Set(reflect.ValueOf(v2))
}

func main() {
	numbers := []int{10, 8, 2, 5, 1, 3, 4, 9, 7, 6}
	sort.Sort(sort.IntSlice(numbers))
	fmt.Println(numbers)
	// sort.Ints(numbers)でも同じことができる
	// Int、Float64、Stringは定義済み。それ以外はソート処理用の構造体を定義する必要がある。

	// 整数の配列
	l1 := []int{10, 8, 2, 5, 1, 3, 4, 9, 7, 6}
	sort.Sort(NewSortwrap(l1, func (i, j int) bool {
		// 配列l1のi番目とj番目の要素を比較する関数
		return l1[i] < l1[j]
	},))
	fmt.Println(l1)
}
