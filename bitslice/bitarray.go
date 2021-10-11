package bitslice

import (
	"bytes"
	"encoding/binary"
	"io"
)

type BitSlice struct {
	Slice     []bool
	ByteOrder binary.ByteOrder
}

func NewBitSliceFromBool(b []bool, bo binary.ByteOrder) *BitSlice {
	return &BitSlice{
		b, bo,
	}
}

func NewBitSliceFromBytes(b []byte, bo binary.ByteOrder) (*BitSlice, error) {
	buf := bytes.NewBuffer(b)
	return NewBitSliceFromReader(buf, bo, uint(len(b)))
}

func NewBitSliceFromReader(r io.Reader, bo binary.ByteOrder, bytesSize uint) (*BitSlice, error) {
	var inBytes = make([]byte, bytesSize)
	err := binary.Read(r, bo, &inBytes)
	if err != nil {
		return nil, err
	}
	bs := &BitSlice{
		Slice:     make([]bool, bytesSize*8),
		ByteOrder: bo,
	}

	lFunc := func(bPtr *byte, index int) {
		bs.Slice[index] = (*bPtr<<7)>>7 == 1
		*bPtr = *bPtr >> 1
	}

	bFunc := func(bPtr *byte, index int) {
		bs.Slice[index] = *bPtr>>7 == 1
		*bPtr = *bPtr << 1
	}
	nowFunc := lFunc
	if bo.String() == binary.BigEndian.String() {
		nowFunc = bFunc
	}

	for i, b := range inBytes {
		for j, _ := range [8]bool{} {
			nowFunc(&b, 8*i+j)
		}
	}
	return bs, nil
}

func (s BitSlice) ToBytes() []byte {
	var packed []byte
	var flagTrue byte
	var flagTrueDefault byte = 1

	lFunc := func() {
		flagTrue = flagTrue << 1
	}

	bFunc := func() {
		flagTrue = flagTrue >> 1
	}
	nowFunc := lFunc
	if s.ByteOrder.String() == binary.BigEndian.String() {
		nowFunc = bFunc
		flagTrueDefault = flagTrueDefault << 7
	}
	flagTrue = flagTrueDefault

	for i, flag := range s.Slice {
		if i%8 == 0 {
			packed = append(packed, 0)
			flagTrue = flagTrueDefault
		}
		if flag {
			packed[i/8] |= flagTrue
		}
		nowFunc()
	}
	return packed
}

func (s BitSlice) ToBuffer(w io.Writer) error {
	err := binary.Write(w, s.ByteOrder, s.ToBytes())
	if err != nil {
		return err
	}
	return nil
}
