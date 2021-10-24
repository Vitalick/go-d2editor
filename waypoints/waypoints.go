package waypoints

import (
	"bytes"
	"encoding/binary"
	"errors"
	"io"
)

const (
	defaultHeaderString = "WS"
	headerLength        = 2
	magicLength         = 6
)

var (
	defaultHeader = [headerLength]byte{}
	defaultMagic  = [magicLength]byte{6, 0, 0, 0, 42, 1}
	wrongHeader   = errors.New("wrong waypoints header")
)

func init() {
	copy(defaultHeader[:], defaultHeaderString[:])
}

//Waypoints ...
type Waypoints struct {
	header    [headerLength]byte
	magic     [magicLength]byte
	Normal    Difficulty `json:"normal"`
	Nightmare Difficulty `json:"nightmare"`
	Hell      Difficulty `json:"hell"`
}

//NewEmptyWaypoints returns empty Waypoints
func NewEmptyWaypoints() (*Waypoints, error) {
	q := &Waypoints{}
	q.header = defaultHeader
	q.magic = defaultMagic
	var err error
	q.Normal, err = NewEmptyDifficulty()
	if err != nil {
		return nil, err
	}
	q.Nightmare, err = NewEmptyDifficulty()
	if err != nil {
		return nil, err
	}
	q.Hell, err = NewEmptyDifficulty()
	if err != nil {
		return nil, err
	}
	return q, nil
}

//NewWaypoints returns Waypoints from packed bytes
func NewWaypoints(r io.Reader) (*Waypoints, error) {
	q := &Waypoints{}

	if err := binary.Read(r, binaryEndian, &q.header); err != nil {
		return nil, err
	}
	headerString := string(bytes.Trim(q.header[:], "\x00"))
	if headerString != defaultHeaderString {
		return nil, wrongHeader
	}
	if err := binary.Read(r, binaryEndian, &q.magic); err != nil {
		return nil, err
	}
	if q.magic != defaultMagic {
		q.magic = defaultMagic
	}
	var err error
	q.Normal, err = NewDifficulty(r)
	if err != nil {
		return nil, err
	}
	q.Nightmare, err = NewDifficulty(r)
	if err != nil {
		return nil, err
	}
	q.Hell, err = NewDifficulty(r)
	if err != nil {
		return nil, err
	}
	return q, nil
}

//GetPacked returns packed Waypoints into []byte
func (q *Waypoints) GetPacked() ([]byte, error) {
	var buf bytes.Buffer
	var packedAct []byte
	var err error

	if err = binary.Write(&buf, binaryEndian, q.header); err != nil {
		return nil, err
	}
	if err = binary.Write(&buf, binaryEndian, q.magic); err != nil {
		return nil, err
	}

	packedAct, err = q.Normal.GetPacked()
	if err != nil {
		return nil, err
	}
	if err = binary.Write(&buf, binaryEndian, packedAct); err != nil {
		return nil, err
	}

	packedAct, err = q.Nightmare.GetPacked()
	if err != nil {
		return nil, err
	}
	if err = binary.Write(&buf, binaryEndian, packedAct); err != nil {
		return nil, err
	}

	packedAct, err = q.Hell.GetPacked()
	if err != nil {
		return nil, err
	}
	if err = binary.Write(&buf, binaryEndian, packedAct); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
