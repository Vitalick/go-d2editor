package quests

import (
	"io"
)

type Act struct {
	Id     ActId
	Quests []Quest
}

//NewAct returns Act from packed bytes
func NewAct(r io.Reader, actId ActId) (Act, error) {

	a := Act{Id: actId}
	count := a.QuestCount()
	for j := 0; byte(j) < count; j++ {
		q, err := NewQuest(r)
		if err != nil {
			return a, err
		}
		a.Quests = append(a.Quests, q)
	}

	return a, nil
}

//QuestCount returns quantity of Quest in current Act
func (a *Act) QuestCount() byte {
	return actCountTable[a.Id]
}

//GetPacked returns packed Act into []uint16
func (a *Act) GetPacked() []uint16 {
	var out []uint16
	count := a.QuestCount()
	for j := 0; byte(j) < count; j++ {
		out = append(out, a.Quests[j].GetPacked())
	}
	return out
}
