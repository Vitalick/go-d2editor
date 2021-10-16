package quests

import (
	"encoding/json"
	"errors"
	"github.com/vitalick/d2s/bitslice"
	"github.com/vitalick/d2s/consts"
	"github.com/vitalick/d2s/utils"
	"io"
)

type QuestImportExport map[string]bool

var questNotExists = errors.New("quest not exist")

//Quest ...
type Quest struct {
	actID   consts.ActID
	questID ActQuest
	flags   []bool
}

//NewEmptyQuest returns empty Quest
func NewEmptyQuest(a consts.ActID, qID ActQuest) (*Quest, error) {
	actMap, ok := actQuestsMap[a]
	if !ok {
		return nil, actNotExists
	}
	_, ok = actMap[qID]
	if !ok {
		return nil, questNotExists
	}
	return &Quest{actID: a, questID: qID, flags: make([]bool, questFlagCount)}, nil
}

//NewQuest returns Quest from packed bytes
func NewQuest(r io.Reader, a consts.ActID, qID ActQuest) (*Quest, error) {
	q, err := NewEmptyQuest(a, qID)
	if err != nil {
		return nil, err
	}
	bs, err := bitslice.NewBitSliceFromReader(r, binaryEndian, 2)
	if err != nil {
		return nil, err
	}
	q.flags = bs.Slice[:questFlagCount]
	return q, nil
}

func (q *Quest) String() string {
	return actQuestsMap[q.actID][q.questID]
}

//GetFlag ...
func (q *Quest) GetFlag(flag QuestFlag) bool {
	return q.flags[flag]
}

//SetFlag ...
func (q *Quest) SetFlag(flag QuestFlag, val bool) {
	if len(q.flags) < questFlagCount {
		oldFlags := q.flags
		q.flags = make([]bool, questFlagCount)
		for i, f := range oldFlags {
			q.flags[i] = f
		}
	}
	q.flags[flag] = val
}

// ExportMap ...
func (q *Quest) ExportMap() *QuestImportExport {
	exportMap := QuestImportExport{}
	for flag := range make([]bool, questFlagCount) {
		qf := QuestFlag(flag)
		exportMap[utils.TitleToJSONTitle(qf.String())] = q.GetFlag(qf)
	}
	return &exportMap
}

// MarshalJSON ...
func (q *Quest) MarshalJSON() ([]byte, error) {
	return json.Marshal(q.ExportMap())
}

// ImportMap ...
func (q *Quest) ImportMap(importMap QuestImportExport) {
	for flag := range make([]bool, questFlagCount) {
		qf := QuestFlag(flag)
		flagStatus, ok := importMap[utils.TitleToJSONTitle(qf.String())]
		if ok {
			q.SetFlag(qf, flagStatus)
		}
	}
}

// UnmarshalJSON ...
func (q *Quest) UnmarshalJSON(data []byte) error {
	importMap := QuestImportExport{}
	if err := json.Unmarshal(data, &importMap); err != nil {
		return err
	}
	q.ImportMap(importMap)
	return nil
}

//GetPacked returns packed Quest into []byte
func (q *Quest) GetPacked() []byte {
	bs := bitslice.NewBitSliceFromBool(q.flags, binaryEndian)
	return bs.ToBytes()
}
