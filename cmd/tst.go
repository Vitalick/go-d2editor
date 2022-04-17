package main

import (
	"fmt"
	"github.com/vitalick/go-d2editor/bitworker"
	"math"
)

func main() {
	b := []byte{7, 15}
	br := bitworker.NewBitReader(b)
	bw := bitworker.NewBitWriter(b)
	fmt.Println(br)
	res, err := br.ReadBits(4, 9)
	if err != nil {
		panic(err)
	}
	fmt.Println(res)
	if err := bw.WriteBits(math.MaxUint32, 4, 9); err != nil {
		panic(err)
	}
	fmt.Println(bw)
	fmt.Println(br)
}
