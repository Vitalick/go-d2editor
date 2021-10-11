package quests

import (
	"github.com/vitalick/d2s/bitslice"
	"io"
)

const (
	questMove = questFlagCount - 1
)

type Quest [questFlagCount]bool

//NewQuest returns Quest from packed bytes
func NewQuest(r io.Reader) (Quest, error) {

	q := Quest{}
	bs, err := bitslice.NewBitSliceFromReader(r, binaryEndian, 2)
	if err != nil {
		return q, err
	}
	//fmt.Println(bs.Slice)
	copy(q[:], bs.Slice[:questFlagCount])
	//fmt.Println(q)

	return q, nil
}

//GetPacked returns packed Quest into []byte
func (q *Quest) GetPacked() []byte {
	bs := bitslice.NewBitSliceFromBool(q[:], binaryEndian)
	return bs.ToBytes()
}
