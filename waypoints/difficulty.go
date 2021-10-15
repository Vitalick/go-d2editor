package waypoints

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"github.com/vitalick/d2s/bitslice"
	"github.com/vitalick/d2s/consts"
	"github.com/vitalick/d2s/utils"
	"io"
)

//Difficulty ...
type Difficulty struct {
	header        [2]byte
	actsWaypoints []bool
	magic         [17]byte
}

//NewDifficulty returns Difficulty from packed bytes
func NewDifficulty(r io.Reader) (Difficulty, error) {
	d := Difficulty{}

	if err := binary.Read(r, binaryEndian, &d.header); err != nil {
		return d, err
	}
	bs, err := bitslice.NewBitSliceFromReader(r, binaryEndian, 5)
	if err != nil {
		return d, err
	}
	d.actsWaypoints = bs.Slice
	err = binary.Read(r, binaryEndian, &d.magic)
	if err != nil {
		return d, err
	}

	return d, nil
}

//GetWaypointState ...
func (d *Difficulty) GetWaypointState(w ActWaypoint) bool {
	return d.actsWaypoints[w]
}

//SetWaypointState ...
func (d *Difficulty) SetWaypointState(w ActWaypoint, val bool) {
	d.actsWaypoints[w] = val
}

//GetActWaypoints ...
func (d *Difficulty) GetActWaypoints(a consts.ActID) []ActWaypoint {
	return actWaypointsMap[a]
}

// MarshalJSON ...
func (d *Difficulty) MarshalJSON() ([]byte, error) {
	exportMap := map[string]map[string]bool{}
	for act, waypoints := range actWaypointsMap {
		actMap := map[string]bool{}
		for _, wp := range waypoints {
			actMap[utils.TitleToJSONTitle(wp.String())] = d.actsWaypoints[wp]
		}
		exportMap[utils.TitleToJSONTitle(act.String())] = actMap
	}
	return json.Marshal(&exportMap)
}

//GetPacked returns packed Difficulty into []byte
func (d *Difficulty) GetPacked() ([]byte, error) {
	var buf bytes.Buffer
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
