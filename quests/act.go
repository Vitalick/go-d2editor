package quests

import (
	"encoding/json"
	"github.com/vitalick/d2s/consts"
	"github.com/vitalick/d2s/utils"
	"io"
)

type Act struct {
	id     consts.ActId
	quests []Quest
}

//NewAct returns Act from packed bytes
func NewAct(r io.Reader, actId consts.ActId) (Act, error) {

	a := Act{id: actId}
	count := a.QuestCount()
	for j := 0; j < count; j++ {
		q, err := NewQuest(r, actId, ActQuest(j))
		if err != nil {
			return a, err
		}
		a.quests = append(a.quests, *q)
	}

	return a, nil
}

//QuestCount returns quantity of Quest in current Act
func (a *Act) QuestCount() int {
	return actLengths[a.id]
}

//GetQuest returns Quest in current Act
func (a *Act) GetQuest(q ActQuest) *Quest {
	return &a.quests[q]
}

//String ...
func (a *Act) String() string {
	return a.id.String()
}

// ExportMap ...
func (a *Act) ExportMap() *map[string]interface{} {
	exportMap := map[string]interface{}{}
	for _, q := range a.quests {
		exportMap[utils.TitleToJsonTitle(q.String())] = q.ExportMap()
	}
	return &exportMap
}

// MarshalJSON ...
func (a *Act) MarshalJSON() ([]byte, error) {
	return json.Marshal(a.ExportMap)
}

//GetPacked returns packed Act into []byte
func (a *Act) GetPacked() []byte {
	var out []byte
	count := a.QuestCount()
	for j := 0; j < count; j++ {
		out = append(out, a.quests[j].GetPacked()...)
	}
	return out
}
