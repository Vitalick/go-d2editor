package npcdialogs

import (
	"encoding/json"
	"errors"
	"github.com/vitalick/bitslice"
	"github.com/vitalick/go-d2editor/utils"
)

type DifficultyImportMap map[string]NPCDialogData

var (
	bitSliceSizeError = errors.New("wrong Length of bit slice")
)

//Difficulty ...
type Difficulty []NPCDialogData

//NewEmptyDifficulty returns empty Difficulty
func NewEmptyDifficulty() Difficulty {
	d := make(Difficulty, npcDialogsCount)
	return d
}

//NewDifficulty returns Difficulty from packed bytes
func NewDifficulty(slice []bool, dPosition uint) (Difficulty, error) {
	d := NewEmptyDifficulty()

	lenBytes := len(slice) / 8
	if len(slice)%8 > 0 {
		lenBytes += 1
	}
	if lenBytes != bitSliceSizeBytes {
		return nil, bitSliceSizeError
	}
	innerSlice := slice[int(dPosition)*npcDialogsCount:]
	for i := range d {
		d[i].Introduction = innerSlice[i]
		d[i].Congratulations = innerSlice[i+dialogTypeOffset]
	}
	return d, nil
}

//GetNPCDialogData ...
func (d *Difficulty) GetNPCDialogData(dd NPCDialog) *NPCDialogData {
	return &(*d)[dd]
}

// MarshalJSON ...
func (d *Difficulty) MarshalJSON() ([]byte, error) {
	exportMap := DifficultyImportMap{}
	for i, dd := range *d {
		exportMap[utils.TitleToJSONTitle(NPCDialog(i).String())] = dd
	}
	return json.Marshal(&exportMap)
}

// UnmarshalJSON ...
func (d *Difficulty) UnmarshalJSON(data []byte) error {
	importMap := DifficultyImportMap{}
	if err := json.Unmarshal(data, &importMap); err != nil {
		return err
	}
	for npcDialog := range npcDialogMap {
		title := utils.TitleToJSONTitle(npcDialog.String())
		dd, ok := importMap[title]
		if !ok {
			continue
		}
		(*d)[npcDialog] = dd
	}
	return nil
}

//GetPacked returns packed Difficulty into []byte
func (d Difficulty) GetPacked(dPosition uint) (*bitslice.BitSlice, error) {
	outSlice, err := bitslice.NewBitSliceFromBytes(make([]byte, bitSliceSizeBytes), binaryEndian)
	if err != nil {
		return nil, err
	}
	for i := range d {
		outSlice.Slice[i] = d[i].Introduction
		outSlice.Slice[i+dialogTypeOffset] = d[i].Congratulations
	}
	shiftedSlice := outSlice.ShiftRight(int(dPosition) * npcDialogsCount)
	return &shiftedSlice, nil
}
