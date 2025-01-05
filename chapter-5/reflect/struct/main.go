package main

import (
	"fmt"
	"reflect"
)


type structValue struct {
	X64 int64
	Y64 int64
}

func main() {
	structValue := &structValue{33,22}
	rv := reflect.ValueOf(structValue).Elem()
	rt := rv.Type()
	for i := 0; i < rt.NumField(); i++ {
		ftv := rt.Field(i)
		fmt.Println(ftv)
		// ftv はreflect.StructField型
		// このフィールドについてのかた情報を得る
		fv := rv.Field(i)
		fmt.Println(fv)
		// ftv はreflect.Value型
	}
}
