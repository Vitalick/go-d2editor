package bitworker

import (
	"errors"
	"fmt"
	"unsafe"
)

type BitReader struct {
	bitNext uint64
	b       []byte
}

func NewBitReader(b []byte) *BitReader {
	return &BitReader{
		b: b,
	}
}

func (bt *BitReader) String() string {
	return fmt.Sprintf("%08b", bt.b)
}

func (bt *BitReader) ReadBits(start, size uint64) (uint64, error) {
	indexOfDataElements := int(start / 8)
	offsetOfDataElement := int(start & 7)
	if int(size)+offsetOfDataElement > 64 {
		return 0, errors.New("request too large")
	}

	if int((start+size)/8) >= len(bt.b) && start+size != uint64(len(bt.b)*8) {
		return 0, errors.New("slice too small")
	}
	ptrToDataElement := unsafe.Pointer(&bt.b[indexOfDataElements])
	fullData := *((*uint64)(ptrToDataElement))
	offsetData := fullData >> offsetOfDataElement
	sizeMask := uint64((1 << size) - 1)
	bt.bitNext = start + size
	return offsetData & sizeMask, nil
}

func (bt *BitReader) ReadNextBits(size uint64) (uint64, error) {
	return bt.ReadBits(bt.bitNext, size)
}

func (bt *BitReader) ReadNextBitsShortBool() (bool, error) {
	res, err := bt.ReadBits(bt.bitNext, 1)
	if err != nil {
		return false, err
	}
	return res == 1, nil
}

func (bt *BitReader) ReadNextBitsUint64() (uint64, error) {
	return bt.ReadBits(bt.bitNext, 64)
}

func (bt *BitReader) ReadNextBitsUint32() (uint32, error) {
	res, err := bt.ReadBits(bt.bitNext, 32)
	if err != nil {
		return 0, err
	}
	return uint32(res), nil
}

func (bt *BitReader) ReadNextBitsUint16() (uint16, error) {
	res, err := bt.ReadBits(bt.bitNext, 16)
	if err != nil {
		return 0, err
	}
	return uint16(res), nil
}

func (bt *BitReader) ReadNextBitsByte() (byte, error) {
	res, err := bt.ReadBits(bt.bitNext, 8)
	if err != nil {
		return 0, err
	}
	return byte(res), nil
}

func (bt *BitReader) ReadNextBitsByteSlice(size uint) ([]byte, error) {
	res := make([]byte, size)
	var err error
	for i := range res {
		res[i], err = bt.ReadNextBitsByte()
		if err != nil {
			return nil, err
		}
	}
	return res, nil
}

func (bt *BitReader) ReadNextBitsByteArray(out []byte) error {
	res := make([]byte, len(out))
	var err error
	for i := range res {
		res[i], err = bt.ReadNextBitsByte()
		if err != nil {
			return err
		}
	}
	copy(out, res[:len(out)])
	return nil
}

func (bt *BitReader) Read(p []byte) (int, error) {
	err := bt.ReadNextBitsByteArray(p)
	if err != nil {
		return 0, err
	}
	return len(p), nil
}

func (bt *BitReader) ReadNextBitsShortBoolSlice(size uint) ([]bool, error) {
	res := make([]bool, size)
	var err error
	for i := range res {
		res[i], err = bt.ReadNextBitsShortBool()
		if err != nil {
			return nil, err
		}
	}
	return res, nil
}

func (bt *BitReader) MoveToByte(to int) error {
	n := int(bt.bitNext/8) + to
	if n > len(bt.b) {
		return errors.New("EOF")
	}
	if n < 0 {
		return errors.New("SOF")
	}
	bt.bitNext = uint64(n * 8)
	return nil
}

func (bt *BitReader) MoveToNextByte() error {
	return bt.MoveToByte(1)
}

func (bt *BitReader) MoveToPreviousByte() error {
	return bt.MoveToByte(-1)
}

func (bt *BitReader) MoveToBit(to int) error {
	n := int(bt.bitNext) + to
	if n >= len(bt.b)*8 {
		return errors.New("EOF")
	}
	if n < 0 {
		return errors.New("SOF")
	}
	bt.bitNext = uint64(n)
	return nil
}
