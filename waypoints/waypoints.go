package waypoints

import (
	"bytes"
	"encoding/binary"
	"io"
)

const (
	defaultWaypointsHeader = "WS"
	waypointsHeaderLength  = 2
	waypointsMagicLength   = 6
)

//Waypoints ...
type Waypoints struct {
	header    [waypointsHeaderLength]byte
	magic     [waypointsMagicLength]byte
	Normal    Difficulty `json:"normal"`
	Nightmare Difficulty `json:"nightmare"`
	Hell      Difficulty `json:"hell"`
}

//NewWaypoints returns Waypoints from packed bytes
func NewWaypoints(r io.Reader) (*Waypoints, error) {
	q := &Waypoints{}

	if err := binary.Read(r, binaryEndian, &q.header); err != nil {
		return nil, err
	}
	if err := binary.Read(r, binaryEndian, &q.magic); err != nil {
		return nil, err
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
	headerString := string(bytes.Trim(q.header[:], "\x00"))
	if headerString != defaultWaypointsHeader {
		var charName [waypointsHeaderLength]byte
		copy(charName[:], defaultWaypointsHeader[:])
		q.header = charName
	}
	if q.magic == [waypointsMagicLength]byte{} {
		q.magic = [waypointsMagicLength]byte{6, 0, 0, 0, 42, 1}
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
