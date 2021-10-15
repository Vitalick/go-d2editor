package quests

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"github.com/vitalick/d2s/consts"
	"github.com/vitalick/d2s/utils"
	"io"
)

//Difficulty ...
type Difficulty []Act

//NewDifficulty returns Difficulty from packed bytes
func NewDifficulty(r io.Reader) (Difficulty, error) {
	d := Difficulty{}
	for i := range make([]bool, consts.ActsCount) {
		act, err := NewAct(r, consts.ActID(i))
		if err != nil {
			return d, err
		}
		d = append(d, act)
	}
	return d, nil
}

//GetAct returns Act in current Difficulty
func (d *Difficulty) GetAct(a consts.ActID) *Act {
	return &(*d)[a]
}

//GetQuest returns Quest in current Difficulty
func (d *Difficulty) GetQuest(a consts.ActID, q ActQuest) *Quest {
	return d.GetAct(a).GetQuest(q)
}

// ExportMap ...
func (d Difficulty) ExportMap() *map[string]interface{} {
	exportMap := map[string]interface{}{}
	for _, a := range d {
		exportMap[utils.TitleToJSONTitle(a.String())] = a.ExportMap()
	}
	return &exportMap
}

// MarshalJSON ...
func (d Difficulty) MarshalJSON() ([]byte, error) {
	return json.Marshal(d.ExportMap())
}

//GetPacked returns packed Difficulty into []byte
func (d *Difficulty) GetPacked() ([]byte, error) {
	var buf bytes.Buffer

	for _, act := range *d {
		if err := binary.Write(&buf, binaryEndian, act.GetPacked()); err != nil {
			return nil, err
		}
	}
	return buf.Bytes(), nil
}
