package d2editor

import (
	"github.com/vitalick/go-d2editor/bitworker"
)

const (
	hardcoreByte  = 1 << 2
	deadByte      = 1 << 3
	expansionByte = 1 << 5
	ladderByte    = 1 << 6
)

//Status ...
type Status struct {
	IsHardcore  bool `json:"is_hardcore"`
	IsDead      bool `json:"is_dead"`
	IsExpansion bool `json:"is_expansion"`
	IsLadder    bool `json:"is_ladder"`
}

//NewStatus ...
func NewStatus(br *bitworker.BitReader) (*Status, error) {
	flags, err := br.ReadNextBitsByte()
	if err != nil {
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
