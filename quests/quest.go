package quests

import (
	"encoding/binary"
	"io"
)

const (
	questMove = questFlagCount - 1
)

type Quest [questFlagCount]bool

//NewQuest returns Quest from packed bytes
func NewQuest(r io.Reader) (Quest, error) {

	var b uint16
	q := Quest{}
	if err := binary.Read(r, binaryEndian, &b); err != nil {
		return q, err
	}
	for i := range q {
		q[i] = (b<<questMove)>>questMove == 1
		b = b >> 1
	}

	return q, nil
}

//GetPacked returns packed Quest into one uint16
func (q *Quest) GetPacked() uint16 {
	var packed uint16
	var flagTrue uint16 = 1
	for _, flag := range q {
		if flag {
			packed |= flagTrue
		}
		flagTrue = flagTrue << 1
	}
	return packed
}
