package d2editor

import (
	"encoding/binary"
	"github.com/vitalick/go-d2editor/consts"
	"io"
)

const defaultMagic = 0xAA55AA55

//Header ...
type Header struct {
	Magic    uint32 `json:"magic"`
	Version  uint32 `json:"version"`
	Filesize uint32 `json:"filesize"`
	Checksum uint32 `json:"checksum"`
}

//NewEmptyHeader returns empty Header
func NewEmptyHeader(version uint) *Header {
	return &Header{defaultMagic, uint32(version), 0, 0}
}

//NewHeader returns Header from packed bytes
func NewHeader(r io.Reader) (*Header, error) {
	h := &Header{}
	if err := binary.Read(r, consts.BinaryEndian, h); err != nil {
		return nil, err
	}
	if h.Magic == 0 {
		h.Magic = defaultMagic
	}
	return h, nil
}

//Fix changes filesize and checksum on struct
func (h *Header) Fix(c *Character) error {
	err := h.fixBytes(c)
	if err != nil {
		return err
	}
	return nil
}

func (h *Header) fixBytes(c *Character) error {
	if err := h.fixSize(c); err != nil {
		return err
	}
	if err := h.fixChecksum(c); err != nil {
		return err
	}
	return nil
}

func (h *Header) fixChecksum(c *Character) error {
	h.Checksum = 0
	b, err := c.GetBytes()
	if err != nil {
		return err
	}
	checksum := 0
	for _, b := range b {
		secondVal := 0
		if checksum < 0 {
			secondVal = 1
		}
		checksum = int(b) + checksum*2 + secondVal
	}
	h.Checksum = uint32(checksum)
	return nil
}

func (h *Header) fixSize(c *Character) error {
	b, err := c.GetBytes()
	if err != nil {
		return err
	}
	h.Filesize = uint32(len(b))
	return nil
}
