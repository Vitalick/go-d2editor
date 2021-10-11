package quests

import (
	"github.com/vitalick/d2s/consts"
	"io"
)

type Act struct {
	Id     consts.ActId
	Quests []Quest
}

//NewAct returns Act from packed bytes
func NewAct(r io.Reader, actId consts.ActId) (Act, error) {

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

//GetPacked returns packed Act into []byte
func (a *Act) GetPacked() []byte {
	var out []byte
	count := a.QuestCount()
	for j := 0; byte(j) < count; j++ {
		out = append(out, a.Quests[j].GetPacked()...)
	}
	return out
}
