package main

import (
	"errors"
	"reflect"
)

type Point struct {
	X int
	Y int
}

func main() {
	Marshal(&Point{10, 100})
}

func Marshal(v interface{}) ([]byte, error) {
	rv := reflect.ValueOf(v)
	switch rv.Kind() {
	case reflect.Map:
		// map用のコード
	case reflect.Struct:
		// struct用のコード
	default:
		// それ以外の場合はエラー
		return nil, errors.New("unsupported type (" + rv.Type().String() + ")")
	}
}
