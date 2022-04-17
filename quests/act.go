package quests

import (
	"encoding/json"
	"errors"
	"github.com/vitalick/go-d2editor/bitworker"
	"github.com/vitalick/go-d2editor/consts"
	"github.com/vitalick/go-d2editor/utils"
)

type ActImportMap map[string]QuestImportExport

var actNotExists = errors.New("act not exist")

//Act ...
type Act struct {
	id     consts.ActID
	quests []Quest
}

//NewEmptyAct returns empty Act
func NewEmptyAct(a consts.ActID) (*Act, error) {
	act := &Act{id: a}
	if int(a) >= len(actLengths) {
		return nil, actNotExists
	}
	actLength := actLengths[a]
	act.quests = make([]Quest, actLength)
	for q := range act.quests {
		quest, err := NewEmptyQuest(a, ActQuest(q))
		if err != nil {
			return nil, err
		}
		act.quests[q] = *quest
	}
	return act, nil
}

//NewAct returns Act from packed bytes
func NewAct(br *bitworker.BitReader, actID consts.ActID) (*Act, error) {

	a := &Act{id: actID}
	count := a.QuestCount()
	for j := 0; j < count; j++ {
		q, err := NewQuest(br, actID, ActQuest(j))
		if err != nil {
			return nil, err
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
func (a *Act) ExportMap() *ActImportMap {
	exportMap := ActImportMap{}
	for _, q := range a.quests {
		exportMap[utils.TitleToJSONTitle(q.String())] = *q.ExportMap()
	}
	return &exportMap
}

// MarshalJSON ...
func (a *Act) MarshalJSON() ([]byte, error) {
	return json.Marshal(a.ExportMap)
}

// ImportMap ...
func (a *Act) ImportMap(importMap ActImportMap) error {
	if int(a.id) >= consts.ActsCount {
		return actNotExists
	}
	actMap := actQuestsMap[a.id]
	for quest := range make([]struct{}, actLengths[a.id]) {
		if quest >= actLengths[a.id] {
			continue
		}
		actQuest := ActQuest(quest)
		questMapTitle := actMap[actQuest]
		questMapTitle = utils.TitleToJSONTitle(questMapTitle)
		questImportMap, ok := importMap[questMapTitle]
		if !ok {
			continue
		}
		var questObj *Quest
		if len(a.quests) > int(actQuest) {
			questObj = &a.quests[int(actQuest)]
		}
		if questObj == nil {
			var err error
			questObj, err = NewEmptyQuest(a.id, actQuest)
			if err != nil {
				return err
			}
		}
		questObj.ImportMap(questImportMap)
	}
	return nil
}

// UnmarshalJSON ...
func (a *Act) UnmarshalJSON(data []byte) error {
	importMap := ActImportMap{}
	if err := json.Unmarshal(data, &importMap); err != nil {
		return err
	}
	return a.ImportMap(importMap)
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
