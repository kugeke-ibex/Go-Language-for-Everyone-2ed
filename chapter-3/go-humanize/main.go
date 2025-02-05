package main

import (
	"fmt"
	"github.com/dustin/go-humanize"
	"os"
)

func main () {
	name := os.Args[1]
	s, _ := os.Stat(name)

	fmt.Printf(
		"%s: %s\n",
		name,
		humanize.Bytes(uint64(s.Size())),
	)

	// ./main image.zip
	// image.zip: 4.7 kB
	// if bw, err := humanize.ParseBytes(maxBandWidth);
	// err != nil {
	// 	fmt.Println("Cannnot parse -max-bandwidth", err) // パースできない場合はエラー
	// 	os.Exit(1)
	// } else {
	// 	conf.maxBandWidth
	// }
}
