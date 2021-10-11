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

type Waypoints struct {
	Header    [waypointsHeaderLength]byte `json:"header"`
	Magic     [waypointsMagicLength]byte  `json:"magic"`
	Normal    Difficulty                  `json:"normal"`
	Nightmare Difficulty                  `json:"nightmare"`
	Hell      Difficulty                  `json:"hell"`
}

//NewWaypoints returns Waypoints from packed bytes
func NewWaypoints(r io.Reader) (*Waypoints, error) {
	q := &Waypoints{}

	if err := binary.Read(r, binaryEndian, &q.Header); err != nil {
		return nil, err
	}
	if err := binary.Read(r, binaryEndian, &q.Magic); err != nil {
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
	headerString := string(bytes.Trim(q.Header[:], "\x00"))
	if headerString != defaultWaypointsHeader {
		var charName [waypointsHeaderLength]byte
		copy(charName[:], defaultWaypointsHeader[:])
		q.Header = charName
	}
	if q.Magic == [waypointsMagicLength]byte{} {
		q.Magic = [waypointsMagicLength]byte{6, 0, 0, 0, 42, 1}
	}
	return q, nil
}

//GetPacked returns packed Waypoints into []byte
func (q *Waypoints) GetPacked() ([]byte, error) {
	var buf bytes.Buffer
	var packedAct []byte
	var err error

	if err = binary.Write(&buf, binaryEndian, q.Header); err != nil {
		return nil, err
	}
	if err = binary.Write(&buf, binaryEndian, q.Magic); err != nil {
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
