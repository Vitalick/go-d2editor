package waypoints

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"github.com/vitalick/bitslice"
	"github.com/vitalick/go-d2editor/bitworker"
	"github.com/vitalick/go-d2editor/consts"
	"github.com/vitalick/go-d2editor/utils"
)

type ActImportMap map[string]bool
type DifficultyImportMap map[string]ActImportMap

//Difficulty ...
type Difficulty struct {
	header        [2]byte
	actsWaypoints []bool
	magic         [17]byte
}

//NewEmptyDifficulty returns empty Difficulty
func NewEmptyDifficulty() (Difficulty, error) {
	d := Difficulty{}
	d.header = [2]byte{2, 1}
	d.actsWaypoints = make([]bool, actWaypointsCount)
	return d, nil
}

//NewDifficulty returns Difficulty from packed bytes
func NewDifficulty(br *bitworker.BitReader) (Difficulty, error) {
	d := Difficulty{}
	if err := br.ReadNextBitsByteArray(d.header[:]); err != nil {
		return d, err
	}
	d.header = [2]byte{2, 1}
	res, err := br.ReadNextBitsShortBoolSlice(5 * 8)
	if err != nil {
		return d, err
	}
	d.actsWaypoints = res[:actWaypointsCount]
	if err := br.ReadNextBitsByteArray(d.magic[:]); err != nil {
		return d, err
	}

	return d, nil
}

func (d *Difficulty) updateActsWaypoints() {
	if len(d.actsWaypoints) < actWaypointsCount {
		oldWp := d.actsWaypoints
		d.actsWaypoints = make([]bool, actWaypointsCount)
		for i, wp := range oldWp {
			d.actsWaypoints[i] = wp
		}
	}
}

//GetWaypointState ...
func (d *Difficulty) GetWaypointState(w ActWaypoint) bool {
	d.updateActsWaypoints()
	return d.actsWaypoints[w]
}

//SetWaypointState ...
func (d *Difficulty) SetWaypointState(w ActWaypoint, val bool) {
	d.updateActsWaypoints()
	d.actsWaypoints[w] = val
}

//GetActWaypoints ...
func (d *Difficulty) GetActWaypoints(a consts.ActID) []ActWaypoint {
	return actWaypointsMap[a]
}

// MarshalJSON ...
func (d *Difficulty) MarshalJSON() ([]byte, error) {
	exportMap := DifficultyImportMap{}
	for i, waypoints := range actWaypointsMap {
		act := consts.ActID(i)
		actMap := ActImportMap{}
		for _, wp := range waypoints {
			actMap[utils.TitleToJSONTitle(wp.String())] = d.actsWaypoints[wp]
		}
		exportMap[utils.TitleToJSONTitle(act.String())] = actMap
	}
	return json.Marshal(&exportMap)
}

// UnmarshalJSON ...
func (d *Difficulty) UnmarshalJSON(data []byte) error {
	importMap := DifficultyImportMap{}
	if err := json.Unmarshal(data, &importMap); err != nil {
		return err
	}
	for i, waypoints := range actWaypointsMap {
		act := consts.ActID(i)
		actTitle := utils.TitleToJSONTitle(act.String())
		actMap, ok := importMap[actTitle]
		if !ok {
			continue
		}
		for _, wp := range waypoints {
			wpTitle := utils.TitleToJSONTitle(wp.String())
			waypointState, ok := actMap[wpTitle]
			if !ok {
				continue
			}
			d.SetWaypointState(wp, waypointState)
		}
	}
	return nil
}

//GetPacked returns packed Difficulty into []byte
func (d *Difficulty) GetPacked() ([]byte, error) {
	var buf bytes.Buffer
	d.header = [2]byte{2, 1}
	if err := binary.Write(&buf, binaryEndian, d.header); err != nil {
		return nil, err
	}
	bs := bitslice.NewBitSliceFromBool(d.actsWaypoints, binaryEndian)
	if err := bs.ToBuffer(&buf); err != nil {
		return nil, err
	}

	if err := binary.Write(&buf, binaryEndian, d.magic); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
