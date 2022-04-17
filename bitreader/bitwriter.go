package bitreader

import (
	"errors"
	"fmt"
	"math"
	"unsafe"
)

type BitWriter struct {
	byteNext uint32
	Bytes    []byte
}

func NewBitWriter(b []byte) *BitWriter {
	return &BitWriter{
		Bytes: b,
	}
}

func (bt *BitWriter) String() string {
	return fmt.Sprintf("%08b", bt.Bytes)
}

func (bt *BitWriter) WriteBits(val, start, size uint32) error {
	indexOfDataElements := int(start / 8)
	offsetOfDataElement := int(start & 7)
	if int(size)+offsetOfDataElement > 32 {
		return errors.New("request too large")
	}

	if int((start+size)/8) >= len(bt.Bytes) {
		newB := make([]byte, int((start+size)/8)+1)
		for i := range bt.Bytes {
			newB[i] = bt.Bytes[i]
		}
	}
	ptrToDataElement := (*uint32)(unsafe.Pointer(&bt.Bytes[indexOfDataElements]))
	offsetData := val << offsetOfDataElement
	sizeMask := uint32((1<<size)-1) << offsetOfDataElement
	inVal := offsetData & sizeMask
	*ptrToDataElement &= math.MaxUint32 ^ sizeMask
	*ptrToDataElement |= inVal
	bt.byteNext = start + size
	return nil
}

func (bt *BitWriter) WriteNextBits(val, size uint32) error {
	return bt.WriteBits(val, bt.byteNext, size)
}

func (bt *BitWriter) MoveToNextByte() error {
	n := bt.byteNext/8 + 1
	bt.byteNext = n * 8
	return nil
}

func (bt *BitWriter) MoveToPreviousByte() error {
	n := int(bt.byteNext/8) - 1
	if n < 0 {
		return errors.New("SOF")
	}
	bt.byteNext = uint32(n * 8)
	return nil
}
