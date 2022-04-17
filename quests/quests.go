package quests

import (
	"bytes"
	"encoding/binary"
	"errors"
	"github.com/vitalick/go-d2editor/bitworker"
)

const (
	defaultHeaderString = "Woo!"
	headerLength        = 4
	magicLength         = 6
)

var (
	defaultHeader = [headerLength]byte{}
	defaultMagic  = [magicLength]byte{6, 0, 0, 0, 42, 1}
	wrongHeader   = errors.New("wrong quests header")
)

func init() {
	copy(defaultHeader[:], defaultHeaderString[:])
}

//Quests ...
type Quests struct {
	header    [headerLength]byte
	magic     [magicLength]byte
	Normal    *Difficulty `json:"normal"`
	Nightmare *Difficulty `json:"nightmare"`
	Hell      *Difficulty `json:"hell"`
}

//NewEmptyQuests returns empty Quests
func NewEmptyQuests() (*Quests, error) {
	q := &Quests{}
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

//NewQuests returns Quests from packed bytes
func NewQuests(br *bitworker.BitReader) (*Quests, error) {
	q := &Quests{}
	if err := br.ReadNextBitsByteArray(q.header[:]); err != nil {
		return nil, err
	}
	headerString := string(bytes.Trim(q.header[:], "\x00"))
	if headerString != defaultHeaderString {
		return nil, wrongHeader
	}
	if err := br.ReadNextBitsByteArray(q.magic[:]); err != nil {
		return nil, err
	}
	if q.magic != defaultMagic {
		q.magic = defaultMagic
	}
	var err error
	q.Normal, err = NewDifficulty(br)
	if err != nil {
		return nil, err
	}
	q.Nightmare, err = NewDifficulty(br)
	if err != nil {
		return nil, err
	}
	q.Hell, err = NewDifficulty(br)
	if err != nil {
		return nil, err
	}
	return q, nil
}

//GetPacked returns packed Quests into []byte
func (q *Quests) GetPacked() ([]byte, error) {
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
