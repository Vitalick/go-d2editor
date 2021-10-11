package waypoints

import (
	"bytes"
	"encoding/binary"
	"github.com/vitalick/d2s/bitslice"
	"io"
)

type Difficulty struct {
	Header        [2]byte
	ActsWaypoints []bool
	Magic         [17]byte
}

//NewDifficulty returns Difficulty from packed bytes
func NewDifficulty(r io.Reader) (Difficulty, error) {
	d := Difficulty{}

	if err := binary.Read(r, binaryEndian, &d.Header); err != nil {
		return d, err
	}
	bs, err := bitslice.NewBitSliceFromReader(r, binaryEndian, 5)
	if err != nil {
		return d, err
	}
	d.ActsWaypoints = bs.Slice
	err = binary.Read(r, binaryEndian, &d.Magic)
	if err != nil {
		return d, err
	}

	return d, nil
}

//GetPacked returns packed Difficulty into []byte
func (d *Difficulty) GetPacked() ([]byte, error) {
	var buf bytes.Buffer
	if err := binary.Write(&buf, binaryEndian, d.Header); err != nil {
		return nil, err
	}
	bs := bitslice.NewBitSliceFromBool(d.ActsWaypoints, binaryEndian)
	if err := bs.ToBuffer(&buf); err != nil {
		return nil, err
	}

	if err := binary.Write(&buf, binaryEndian, d.Magic); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
