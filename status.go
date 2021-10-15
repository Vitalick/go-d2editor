package d2s

import (
	"encoding/binary"
	"github.com/vitalick/d2s/consts"
	"io"
)

const (
	firstByte     = 1
	hardcoreByte  = firstByte << 2
	deadByte      = firstByte << 3
	expansionByte = firstByte << 5
	ladderByte    = firstByte << 6
)

//Status ...
type Status struct {
	IsHardcore  bool `json:"is_hardcore"`
	IsDead      bool `json:"is_dead"`
	IsExpansion bool `json:"is_expansion"`
	IsLadder    bool `json:"is_ladder"`
}

//NewStatus ...
func NewStatus(r io.Reader) (*Status, error) {
	var flags byte
	if err := binary.Read(r, consts.BinaryEndian, &flags); err != nil {
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

//GetFlags return all flags packed in one byte
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
