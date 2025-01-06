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

func main() {
	p := &Person{Name: "山田太郎", Age: 33}
	rt := reflect.TypeOf(p)
	rv := reflect.ValueOf(p)

	if rv.Kind() == reflect.Ptr {
		// reflect.Interfaceも同様
		rv = rv.Elem()
	}
	if rt.Kind() == reflect.Ptr {
		// reflect.Interfaceも同様
		rt = rt.Elem()
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
}
