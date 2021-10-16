package d2editor

import "math"

const emptySkillID = math.MaxUint16

//Skill ...
type Skill struct {
	ID uint32 `json:"id"`
}

func NewEmptySkill() *Skill {
	s := &Skill{}
	s.Clear()
	return s
}

func (s *Skill) Clear() {
	s.ID = emptySkillID
}

func (s *Skill) Set(val uint32) {
	s.ID = val
}
