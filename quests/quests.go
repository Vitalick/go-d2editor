package quests

import (
	"bytes"
	"encoding/binary"
	"io"
)

const (
	defaultQuestsHeader = "Woo!"
	questsHeaderLength  = 4
	questsMagicLength   = 6
)

type Quests struct {
	Header    [questsHeaderLength]byte `json:"header"`
	Magic     [questsMagicLength]byte  `json:"magic"`
	Normal    Difficulty               `json:"normal"`
	Nightmare Difficulty               `json:"nightmare"`
	Hell      Difficulty               `json:"hell"`
}

//NewQuests returns Quests from packed bytes
func NewQuests(r io.Reader) (*Quests, error) {
	q := &Quests{}

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
	if headerString != defaultQuestsHeader {
		var charName [questsHeaderLength]byte
		copy(charName[:], defaultQuestsHeader[:])
		q.Header = charName
	}
	if q.Magic == [questsMagicLength]byte{} {
		q.Magic = [questsMagicLength]byte{6, 0, 0, 0, 42, 1}
	}
	return q, nil
}

//GetPacked returns packed Quests into []byte
func (q *Quests) GetPacked() ([]byte, error) {
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