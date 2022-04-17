package bitworker

import (
	"errors"
	"fmt"
	"math"
	"unsafe"
)

type BitWriter struct {
	bitNext uint64
	Bytes   []byte
}

func NewBitWriter(b []byte) *BitWriter {
	return &BitWriter{
		Bytes: b,
	}
}

func (bt *BitWriter) String() string {
	return fmt.Sprintf("%08b", bt.Bytes)
}

func (bt *BitWriter) WriteBits(val, start, size uint64) error {
	if bt.Bytes == nil {
		bt.Bytes = []byte{}
	}
	indexOfDataElements := int(start / 8)
	offsetOfDataElement := int(start & 7)
	if int(size)+offsetOfDataElement > 64 {
		return errors.New("request too large")
	}

	if int((start+size)/8) >= len(bt.Bytes) && start+size != uint64(len(bt.Bytes)*8) {
		newB := make([]byte, int((start+size)/8)+1)
		for i := range bt.Bytes {
			newB[i] = bt.Bytes[i]
		}
		bt.Bytes = newB
	}
	ptrToDataElement := (*uint64)(unsafe.Pointer(&bt.Bytes[indexOfDataElements]))
	offsetData := val << offsetOfDataElement
	sizeMask := uint64((1<<size)-1) << offsetOfDataElement
	inVal := offsetData & sizeMask
	*ptrToDataElement &= math.MaxUint32 ^ sizeMask
	*ptrToDataElement |= inVal
	bt.bitNext = start + size
	return nil
}

func (bt *BitWriter) WriteNextBits(val, size uint64) error {
	return bt.WriteBits(val, bt.bitNext, size)
}

func (bt *BitWriter) WriteNextBitsShortBool(val bool) error {
	var send uint64
	if val {
		send = 1
	}
	return bt.WriteBits(send, bt.bitNext, 1)
}

func (bt *BitWriter) WriteNextBitsUint64(val uint64) error {
	return bt.WriteBits(val, bt.bitNext, 64)
}

func (bt *BitWriter) WriteNextBitsUint32(val uint32) error {
	return bt.WriteBits(uint64(val), bt.bitNext, 32)
}

func (bt *BitWriter) WriteNextBitsUint16(val uint16) error {
	return bt.WriteBits(uint64(val), bt.bitNext, 16)
}

func (bt *BitWriter) WriteNextBitsByte(val byte) error {
	return bt.WriteBits(uint64(val), bt.bitNext, 8)
}

func (bt *BitWriter) WriteNextBitsByteSlice(val []byte) error {
	for _, v := range val {
		if err := bt.WriteNextBitsByte(v); err != nil {
			return err
		}
	}
	return nil
}

func (bt *BitWriter) WriteNextBitsBoolSlice(val []bool) error {
	for _, v := range val {
		if err := bt.WriteNextBitsShortBool(v); err != nil {
			return err
		}
	}
	return nil
}

func (bt *BitWriter) WriteNextBitsByteArray(val []byte) error {
	for _, v := range val {
		if err := bt.WriteNextBitsByte(v); err != nil {
			return err
		}
	}
	return nil
}

func (bt *BitWriter) MoveToByte(to int) error {
	n := int(bt.bitNext/8) + to
	if n < 0 {
		return errors.New("SOF")
	}
	bt.bitNext = uint64(n * 8)
	return nil
}

func (bt *BitWriter) MoveToNextByte() error {
	return bt.MoveToByte(1)
}

func (bt *BitWriter) MoveToPreviousByte() error {
	return bt.MoveToByte(-1)
}

func (bt *BitWriter) MoveToBit(to int) error {
	n := int(bt.bitNext) + to
	if n < 0 {
		return errors.New("SOF")
	}
	bt.bitNext = uint64(n)
	return nil
}
