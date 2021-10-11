package d2s

import (
	"encoding/binary"
	"github.com/vitalick/d2s/consts"
	"io"
)

type Skill struct {
	Id uint32 `json:"id"`
}

//NewSkill ...
func NewSkill(r io.Reader) (*Skill, error) {
	s := &Skill{}
	if err := binary.Read(r, consts.BinaryEndian, &s.Id); err != nil {
		return nil, err
	}
	return s, nil
}
