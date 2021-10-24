package attributes

import (
	"fmt"
	"strings"
)

type bitArray struct {
	s string
	r uint
}

func newBitArray(b []byte) bitArray {
	lenB := len(b)
	newNumbers := make([]byte, lenB)
	for i, j := 0, lenB-1; i <= j; i, j = i+1, j-1 {
		newNumbers[i], newNumbers[j] = b[j], b[i]
	}
	s := strings.Replace(fmt.Sprintf("%08b", newNumbers), " ", "", -1)
	s = s[1 : len(s)-1]
	fmt.Println("newBitArray", len(s))
	fmt.Println(s)
	return bitArray{s: s}
}

func (b *bitArray) GetFirst(amount uint) (uint64, error) {
	lenS := len(b.s)
	s := b.s[lenS-int(amount)-int(b.r) : lenS-int(b.r)]
	//fmt.Println("getFirst", amount)
	//fmt.Println("s", s)
	//fmt.Println()
	var u uint64
	_, err := fmt.Sscanf(s, "%b", &u)
	if err != nil {
		return 0, err
	}
	b.r += amount
	return u, nil
}

func (b bitArray) GetBytes() ([]byte, error) {
	var bytes []byte
	offset := 0
	lenB := len(b.s)
	for offset < lenB {
		end := offset + 8
		if end > lenB {
			end = lenB
		}

		s := b.s[offset:end]
		var bt byte
		_, err := fmt.Sscanf(s, "%b", &bt)
		if err != nil {
			return nil, err
		}
		bytes = append(bytes, bt)
		offset = end
	}
	lenB = len(bytes)
	newNumbers := make([]byte, lenB)
	for i, j := 0, lenB-1; i <= j; i, j = i+1, j-1 {
		newNumbers[i], newNumbers[j] = bytes[j], bytes[i]
	}
	return newNumbers, nil
}
