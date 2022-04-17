package d2editor

import (
	"github.com/vitalick/go-d2editor/bitworker"
)

const defaultMagic = 0xAA55AA55

//Header ...
type Header struct {
	Magic    uint32 `json:"-"`
	Version  uint32 `json:"version"`
	Filesize uint32 `json:"filesize"`
	Checksum uint32 `json:"checksum"`
}

//NewEmptyHeader returns empty Header
func NewEmptyHeader(version uint) *Header {
	return &Header{defaultMagic, uint32(version), 0, 0}
}

//NewHeader returns Header from packed bytes
func NewHeader(br *bitworker.BitReader) (*Header, error) {
	h := &Header{}
	var err error
	h.Magic, err = br.ReadNextBitsUint32()
	if err != nil {
		return nil, err
	}
	h.Version, err = br.ReadNextBitsUint32()
	if err != nil {
		return nil, err
	}
	h.Filesize, err = br.ReadNextBitsUint32()
	if err != nil {
		return nil, err
	}
	h.Checksum, err = br.ReadNextBitsUint32()
	if err != nil {
		return nil, err
	}
	if h.Magic == 0 {
		h.Magic = defaultMagic
	}
	return h, nil
}

func ChecksumAppend(b byte, c int) int {
	secondVal := 0
	if c < 0 {
		secondVal = 1
	}
	c = int(b) + c*2 + secondVal
	return c
}
