package main

import (
	"fmt"
	"sync"
)


// ハッシュマップをなるべく使わずに、structでtypeを定義する
type data struct {
	foo string
	baz string
}

// mapに対する操作はスレッドセーフではないので、複数のgoroutineから同時にアクセスした場合に、変な値を読み込んだり、プログラムがクラッシュしたりする可能性があるので、syncパッケージのRW Mutexを使い、排他制御を加える。

// KeyValue のための型。内部にmapを保持している
type KeyValue struct {
	store map[string]string // key-valueを格納するためのmap
	mu    sync.RWMutex      // 排他制御のためのmutex
}

// KeyValue のための型。内部にmapを保持している
func NewKeyValue() *KeyValue {
	return &KeyValue{store: make(map[string]string)}
}

func (kv *KeyValue) Set(key, val string) {
	kv.mu.Lock()         // まずLock
	defer kv.mu.Unlock() // メソッドを抜けた際にUnlock
	kv.store[key] = val
}

func (kv *KeyValue) Get(key string) (string, bool) {
	kv.mu.RLock()         // 参照用のRLock
	defer kv.mu.RUnlock() // メソッドを抜けた際にRUnlock
	val, ok := kv.store[key]
	return val, ok
}

func main() {
	d := map[string]string{
		"foo": "bar",
		"baz": "qux",
	}

	fmt.Println(d)

	dstruct := data{
		foo: "bar",
		baz: "qux",
	}
	fmt.Println(dstruct)

	kv := NewKeyValue()
	kv.Set("key", "value")
	value, ok := kv.Get("key")
	if ok {
			fmt.Println(value)
	}
}


