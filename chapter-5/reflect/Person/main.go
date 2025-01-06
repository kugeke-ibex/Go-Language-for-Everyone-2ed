package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	Name string
	Age int
}

func (p *Person) setName(name string) {
	p.Name = name
}

func (p Person) printName() string {
	return p.Name
}

type Point struct {
	X int
	Y int
}

func main() {
	p := &Person{Name: "山田太郎", Age: 33}
	rt := reflect.TypeOf(p)
	rv := reflect.ValueOf(p)

	switch rt.Kind() {
	case reflect.Interface, reflect.Ptr:
		rt = rt.Elem()
	}
	switch rv.Kind() {
	case reflect.Interface, reflect.Ptr:
		rv = rv.Elem()
	}

	// fmt.Println(reflect.TypeOf(*p).Method(0))
	for i := 0; i < rv.NumField(); i++ {
		ft := rt.Field(i)
		// i-th目の要素の型(reflect.StructField)
		fv := rv.Field(i)
		// i-th目の要素の値(reflect.Value)
		if ft.PkgPath == "" {
			fmt.Printf("ft(%d) -> %v\n", i, ft)
			fmt.Printf("fv(%d) -> %v\n", i, fv.Interface())
		}
	}

	point := Point{X:10, Y:5}
	rvp := reflect.ValueOf(&point).Elem()
	fmt.Println(rvp)
	if !rvp.Field(0).CanSet() {
		// エラー処理
	}
	if f := rvp.Field(0); f.CanSet() {
		f.SetInt(100)
	}
	fmt.Println(rvp)
}
