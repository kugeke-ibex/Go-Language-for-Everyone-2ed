package main

import (
	"encoding/json"
	"fmt"
)

type target struct {
	Name      string `json:'name'`
	Threshold int    `json:'threshold'`
}

type config struct {
	Addr   string   `json:'addr'`
	Target []target `json:'target'`
}

func main() {
	cfg := config{":8080", []target{{"foo", 3}, {"bar", 4}}}
	b, _ := json.Marshal(&cfg)
	fmt.Println(string(b))

	// インデント付きで出力
	b2, _ := json.MarshalIndent(&cfg, "", " ")
	fmt.Println(string(b2))
}
