package character

import (
	"encoding/binary"
	"io"
)

const (
	firstByte     = 1
	hardcoreByte  = firstByte << 2
	deadByte      = firstByte << 3
	expansionByte = firstByte << 5
	ladderByte    = firstByte << 6
)

type Status struct {
	IsHardcore, IsDead, IsExpansion, IsLadder bool
}

//NewStatus Like Read func
func NewStatus(r io.Reader) (*Status, error) {
	var flags byte
	if err := binary.Read(r, binaryEndian, &flags); err != nil {
		return nil, err
	}
	s := &Status{
		IsHardcore:  hardcoreByte&flags > 0,
		IsDead:      deadByte&flags > 0,
		IsExpansion: expansionByte&flags > 0,
		IsLadder:    ladderByte&flags > 0,
	}
	return s, nil
}

func (s *Status) GetFlags() byte {
	var flags byte
	if s.IsHardcore {
		flags = flags | hardcoreByte
	}
	if s.IsDead {
		flags = flags | deadByte
	}
	if s.IsExpansion {
		flags = flags | expansionByte
	}
	if s.IsLadder {
		flags = flags | ladderByte
	}
	return flags
}
