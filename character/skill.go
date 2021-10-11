package character

import (
	"encoding/binary"
	"io"
)

type Skill struct {
	Id uint32 `json:"id"`
}

//NewSkill ...
func NewSkill(r io.Reader) (*Skill, error) {
	s := &Skill{}
	if err := binary.Read(r, binaryEndian, &s.Id); err != nil {
		return nil, err
	}
	return s, nil
}
