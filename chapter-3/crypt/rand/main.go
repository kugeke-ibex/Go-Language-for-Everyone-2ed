package main

import (
	"encoding/binary"
	crand "crypto/rand"
	"math/rand"
	"time"
)



func main () {
	var s int64
	if err := binary.Read(crand.Reader, binary.LittleEndian, &s); err != nil {
		// crypto/randからReadできなかった場合の代替手段
		s = time.Now().UnixNano()
	}
	rand.Seed(s)
}








