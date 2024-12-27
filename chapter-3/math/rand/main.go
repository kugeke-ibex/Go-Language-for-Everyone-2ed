package main

import (
	"fmt"
	"math/rand"
	// "time"
)

func main() {
	rand.Seed(92)
	// 0 <= n < 100 となるintの乱数を取得
	n := rand.Intn(100)
	fmt.Println(n)
}
