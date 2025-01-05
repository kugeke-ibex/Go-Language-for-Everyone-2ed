package main

import (
	"reflect"
	"fmt"
)

type Mytype struct {

}

func main () {
	rv1 := reflect.ValueOf(1)
	// rv2 := reflect.ValueOf("Hello World")
	// rv3 := reflect.ValueOf([]byte{0xa,0xd})
	// rv4 := reflect.ValueOf(make(chan struct{}))

	fmt.Printf("rv's Int() is %v\n", rv1.Int())
	// fmt.Printf("rv's Int() is %v\n", rv2.Int()) // panic
	// fmt.Printf("rv's Int() is %v\n", rv3.Int()) // panic
	// fmt.Printf("rv's Int() is %v\n", rv4.Int()) // panic

	rvmap := reflect.ValueOf(map[string]int{"foo": 1})
	value := rvmap.MapIndex(reflect.ValueOf("foo"))
	fmt.Println(value)
	// reflect.ValueOf(1)が返ってくる
	rvmap.SetMapIndex(reflect.ValueOf("foo"), reflect.ValueOf(2))
	// "foo": 2と代入
	fmt.Println(rvmap)
	// rvmap.Int() // panic

	var num int64
	if rv1.Kind() == reflect.Int {
		num = rv1.Int()
		fmt.Printf("num is %d\n", num)
	}
	// これは存在しない
	// reflect.Type("MyType")
	// 必ずこのように値から取得する
	fmt.Println(reflect.TypeOf(Mytype{}))
}
