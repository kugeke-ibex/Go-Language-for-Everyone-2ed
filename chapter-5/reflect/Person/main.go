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
		fv := rv.Field(i)
		if ft.PkgPath == "" {
			fmt.Printf("ft(%d) -> %v\n", i, ft)
			fmt.Printf("fv(%d) -> %v\n", i, fv.Interface())
		}
	}
}
