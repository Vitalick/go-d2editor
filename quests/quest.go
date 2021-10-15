package quests

import (
	"encoding/json"
	"github.com/vitalick/d2s/bitslice"
	"github.com/vitalick/d2s/consts"
	"github.com/vitalick/d2s/utils"
	"io"
)

//Quest ...
type Quest struct {
	actID   consts.ActID
	questID ActQuest
	flags   []bool
}

//NewQuest returns Quest from packed bytes
func NewQuest(r io.Reader, a consts.ActID, qID ActQuest) (*Quest, error) {
	q := &Quest{actID: a, questID: qID, flags: make([]bool, questFlagCount)}
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
	q.flags[flag] = val
}

// ExportMap ...
func (q *Quest) ExportMap() *map[string]bool {
	exportMap := map[string]bool{}
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

//GetPacked returns packed Quest into []byte
func (q *Quest) GetPacked() []byte {
	bs := bitslice.NewBitSliceFromBool(q.flags, binaryEndian)
	return bs.ToBytes()
}
