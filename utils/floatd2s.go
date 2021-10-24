package utils

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"github.com/vitalick/bitslice"
	"github.com/vitalick/go-d2editor/consts"
	"io"
	"math"
)

var maxFloatInt = uint32(math.Pow(math.Pow(2, 8), 3))

//FloatD2s type for d2s float
type FloatD2s [4]byte

//GetFloat64 convert FloatD2s to float64
func (f FloatD2s) GetFloat64() float64 {
	bs, err := bitslice.NewBitSliceFromBytes(f[:], binary.BigEndian)
	if err != nil {
		return 0
	}
	buf := bytes.NewBuffer(bs.ShiftRight(8).ToBytes())
	var intPart uint32
	err = binary.Read(buf, binary.BigEndian, &intPart)
	if err != nil {
		return 0
	}
	buf = bytes.NewBuffer(bs.ShiftLeft(24).ShiftRight(24).ToBytes())
	var floatPart uint32
	err = binary.Read(buf, binary.BigEndian, &floatPart)
	if err != nil {
		return 0
	}
	return float64(intPart) + float64(floatPart)/255
}

//SetFloat64 convert float64 to FloatD2s
func (f *FloatD2s) SetFloat64(inFloat float64) error {
	intPart := uint32(inFloat)
	if intPart > maxFloatInt {
		return errors.New(fmt.Sprintf("number should be less or equal then %d", maxFloatInt))
	}
	if intPart == maxFloatInt {
		*f = FloatD2s{255, 255, 255, 255}
		return nil
	}
	floatPart := uint32(math.Ceil((inFloat - float64(intPart)) * 255))
	buf := &bytes.Buffer{}
	if err := binary.Write(buf, binary.BigEndian, intPart); err != nil {
		return err
	}
	bs, err := bitslice.NewBitSliceFromBytes(buf.Bytes(), binary.BigEndian)
	*bs = bs.ShiftLeft(8)
	buf.Reset()
	err = binary.Write(buf, binary.BigEndian, floatPart)
	if err = binary.Write(buf, binary.BigEndian, floatPart); err != nil {
		return err
	}
	bsF, err := bitslice.NewBitSliceFromBytes(buf.Bytes(), binary.BigEndian)
	*bs = bs.Or(*bsF)
	copy(f[:], bs.ToBytes()[:4])
	return nil
}

//FloatD2sGo type for d2s float
type FloatD2sGo float64

func NewFloatD2sGo(r io.Reader) (FloatD2sGo, error) {
	fd2s := FloatD2s{}
	err := binary.Read(r, consts.BinaryEndian, &fd2s)
	if err != nil {
		return 0, err
	}
	f := FloatD2sGo(fd2s.GetFloat64())
	return f, nil
}

//GetPacked convert FloatD2sGo to FloatD2s
func (f FloatD2sGo) GetPacked() (FloatD2s, error) {
	fd2s := FloatD2s{}
	err := fd2s.SetFloat64(float64(f))
	if err != nil {
		return fd2s, err
	}
	return fd2s, nil
}
