package npcdialogs

import (
	"bytes"
	"encoding/binary"
	"errors"
	"github.com/vitalick/bitslice"
	"io"
)

const (
	defaultHeaderString = "w4"
	headerLength        = 2
	defaultLength       = 0x24
)

var (
	defaultHeader = [headerLength]byte{}
	wrongHeader   = errors.New("wrong npc dialogs Header")
)

func init() {
	copy(defaultHeader[:], defaultHeaderString[:])
}

//NPCDialogs ...
type NPCDialogs struct {
	Header    [headerLength]byte
	Length    byte
	Normal    Difficulty `json:"normal"`
	Nightmare Difficulty `json:"nightmare"`
	Hell      Difficulty `json:"hell"`
}

//NewEmptyNPCDialogs returns empty NPCDialogs
func NewEmptyNPCDialogs() *NPCDialogs {
	return &NPCDialogs{
		defaultHeader,
		defaultLength,
		NewEmptyDifficulty(),
		NewEmptyDifficulty(),
		NewEmptyDifficulty(),
	}
}

//NewNPCDialogs returns NPCDialogs from packed bytes
func NewNPCDialogs(r io.Reader) (*NPCDialogs, error) {
	q := &NPCDialogs{}

	if err := binary.Read(r, binaryEndian, &q.Header); err != nil {
		return nil, err
	}
	if err := binary.Read(r, binaryEndian, &q.Length); err != nil {
		return nil, err
	}
	//b := [bitSliceSizeBytes]byte{}
	//
	//if err := binary.Read(r, binaryEndian, &b); err != nil {
	//	return nil, err
	//}
	//fmt.Printf("%0.8b\n", b)
	//return nil, errors.New("test error")
	bitSlice, err := bitslice.NewBitSliceFromReader(r, binaryEndian, bitSliceSizeBytes)
	headerString := string(bytes.Trim(q.Header[:], "\x00"))
	if headerString != defaultHeaderString {
		return nil, wrongHeader
	}
	if q.Length != defaultLength {
		q.Length = defaultLength
	}
	if err != nil {
		return nil, err
	}
	//fmt.Println(bitSlice.Slice)
	q.Normal, err = NewDifficulty(*bitSlice, 0)
	if err != nil {
		return nil, err
	}
	q.Nightmare, err = NewDifficulty(*bitSlice, 1)
	if err != nil {
		return nil, err
	}
	q.Hell, err = NewDifficulty(*bitSlice, 2)
	if err != nil {
		return nil, err
	}
	return q, nil
}

//GetPacked returns packed NPCDialogs into []byte
func (q *NPCDialogs) GetPacked() ([]byte, error) {
	var buf bytes.Buffer

	if err := binary.Write(&buf, binaryEndian, q.Header); err != nil {
		return nil, err
	}
	if err := binary.Write(&buf, binaryEndian, q.Length); err != nil {
		return nil, err
	}

	bsNormal, err := q.Normal.GetPacked(0)
	if err != nil {
		return nil, err
	}

	bsNightmare, err := q.Nightmare.GetPacked(1)
	if err != nil {
		return nil, err
	}

	bsHell, err := q.Hell.GetPacked(2)
	if err != nil {
		return nil, err
	}
	if err := bsNormal.Or(*bsNightmare).Or(*bsHell).ToBuffer(&buf); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
