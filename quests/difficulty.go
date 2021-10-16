package quests

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"github.com/vitalick/d2s/consts"
	"github.com/vitalick/d2s/utils"
	"io"
)

type DifficultyImportMap map[string]ActImportMap

//Difficulty ...
type Difficulty []Act

//NewEmptyDifficulty returns empty Difficulty
func NewEmptyDifficulty() (*Difficulty, error) {
	d := make(Difficulty, consts.ActsCount)
	for i := range d {
		act, err := NewEmptyAct(consts.ActID(i))
		if err != nil {
			return nil, err
		}
		d[i] = *act
	}
	return &d, nil
}

//NewDifficulty returns Difficulty from packed bytes
func NewDifficulty(r io.Reader) (*Difficulty, error) {
	d := make(Difficulty, consts.ActsCount)
	for i := range d {
		act, err := NewAct(r, consts.ActID(i))
		if err != nil {
			return nil, err
		}
		d[i] = *act
	}
	return &d, nil
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
func (d Difficulty) ExportMap() *DifficultyImportMap {
	exportMap := DifficultyImportMap{}
	for _, a := range d {
		exportMap[utils.TitleToJSONTitle(a.String())] = *a.ExportMap()
	}
	return &exportMap
}

// MarshalJSON ...
func (d Difficulty) MarshalJSON() ([]byte, error) {
	return json.Marshal(d.ExportMap())
}

// ImportMap ...
func (d *Difficulty) ImportMap(importMap DifficultyImportMap) error {
	for i, a := range *d {
		actTitle := utils.TitleToJSONTitle(a.String())
		actMap, ok := importMap[actTitle]
		if !ok {
			return actNotExists
		}
		err := (*d)[i].ImportMap(actMap)
		if err != nil {
			return err
		}
	}
	return nil
}

// UnmarshalJSON ...
func (d *Difficulty) UnmarshalJSON(data []byte) error {
	importMap := DifficultyImportMap{}
	if err := json.Unmarshal(data, &importMap); err != nil {
		return err
	}
	return d.ImportMap(importMap)
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
