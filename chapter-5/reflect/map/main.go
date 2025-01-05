package main

import (
	"reflect"
)

func main() {
	rv := reflect.ValueOf(map[string]int{"foo": 1})
	iter := rv.MapRange()
	for iter.Next() {
		key := iter.Key().String()
		switch mapv := iter.Value(); mapv.Kind() {
		case reflect.String:
			value := iter.Value().String()
			// valueが文字列だった場合の処理
		case reflect.Int:
			value := iter.Value().Int()
			// valueが整数だった場合の処理
		}
	}
}
