package bitreader

import (
	"errors"
	"fmt"
	"unsafe"
)

type BitReader struct {
	byteNext uint32
	b        []byte
}

func NewBitReader(b []byte) *BitReader {
	return &BitReader{
		b: b,
	}
}

func (bt *BitReader) String() string {
	return fmt.Sprintf("%08b", bt.b)
}

func (bt *BitReader) ReadBits(start, size uint32) (uint32, error) {
	indexOfDataElements := int(start / 8)
	offsetOfDataElement := int(start & 7)
	if int(size)+offsetOfDataElement > 32 {
		return 0, errors.New("request too large")
	}

	if int((start+size)/8) >= len(bt.b) {
		return 0, errors.New("slice too small")
	}
	ptrToDataElement := unsafe.Pointer(&bt.b[indexOfDataElements])
	fullData := *((*uint32)(ptrToDataElement))
	offsetData := fullData >> offsetOfDataElement
	sizeMask := uint32((1 << size) - 1)
	bt.byteNext = start + size
	return offsetData & sizeMask, nil
}

func (bt *BitReader) ReadNextBits(size uint32) (uint32, error) {
	return bt.ReadBits(bt.byteNext, size)
}

func (bt *BitReader) MoveToNextByte() error {
	n := bt.byteNext/8 + 1
	if int(n) >= len(bt.b) {
		return errors.New("EOF")
	}
	bt.byteNext = n * 8
	return nil
}

func (bt *BitReader) MoveToPreviousByte() error {
	n := int(bt.byteNext/8) - 1
	if n < 0 {
		return errors.New("SOF")
	}
	bt.byteNext = uint32(n * 8)
	return nil
}
