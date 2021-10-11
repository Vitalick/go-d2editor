package d2s

import (
	"encoding/binary"
	"io"
)

const defaultMagic = 0xAA55AA55

type Header struct {
	Magic    uint32 `json:"magic"`
	Version  uint32 `json:"version"`
	Filesize uint32 `json:"filesize"`
	Checksum uint32 `json:"checksum"`
}

//NewHeader Like Read func
func NewHeader(r io.Reader) (*Header, error) {
	h := &Header{}
	if err := binary.Read(r, binaryEndian, h); err != nil {
		return nil, err
	}
	if h.Magic == 0 {
		h.Magic = defaultMagic
	}
	return h, nil
}

//Fix changes filesize and checksum on struct
func (h *Header) Fix(c *Character) error {
	h.Checksum = 0
	b, err := c.GetBytes()
	if err != nil {
		return err
	}
	err = h.fixBytes(b)
	if err != nil {
		return err
	}
	return nil
}

func (h *Header) fixBytes(b []byte) error {
	if err := h.fixChecksum(b); err != nil {
		return err
	}
	if err := h.fixSize(b); err != nil {
		return err
	}
	return nil
}

func (h *Header) fixChecksum(bs []byte) error {
	checksum := 0
	for _, b := range bs {
		secondVal := 0
		if checksum < 0 {
			secondVal = 1
		}
		checksum = int(b) + checksum*2 + secondVal
	}
	h.Checksum = uint32(checksum)
	return nil
}

func (h *Header) fixSize(b []byte) error {
	h.Filesize = uint32(len(b))
	return nil
}
