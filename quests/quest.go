package quests

import (
	"encoding/json"
	"github.com/vitalick/d2s/bitslice"
	"github.com/vitalick/d2s/consts"
	"github.com/vitalick/d2s/utils"
	"io"
)

type Quest struct {
	actId   consts.ActId
	questId ActQuest
	flags   []bool
}

//NewQuest returns Quest from packed bytes
func NewQuest(r io.Reader, a consts.ActId, qId ActQuest) (*Quest, error) {

	q := &Quest{actId: a, questId: qId, flags: make([]bool, questFlagCount)}
	bs, err := bitslice.NewBitSliceFromReader(r, binaryEndian, 2)
	if err != nil {
		return nil, err
	}
	//fmt.Println(bs.Slice)
	q.flags = bs.Slice[:questFlagCount]
	//fmt.Println(q)

	return q, nil
}

func (q *Quest) String() string {
	return actQuestsMap[q.actId][q.questId]
}

func (q *Quest) GetFlag(flag QuestFlag) bool {
	return q.flags[flag]
}

func (q *Quest) SetFlag(flag QuestFlag, val bool) {
	q.flags[flag] = val
}

// ExportMap ...
func (q *Quest) ExportMap() *map[string]bool {
	exportMap := map[string]bool{}
	for flag := range make([]bool, questFlagCount) {
		qf := QuestFlag(flag)
		exportMap[utils.TitleToJsonTitle(qf.String())] = q.GetFlag(qf)
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
