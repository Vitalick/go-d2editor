package d2s

import (
	"encoding/binary"
	"io"
)

const locationsCount = 3

//Locations struct shown opened acts and difficulties information
type Locations struct {
	Normal    Location `json:"normal"`
	Nightmare Location `json:"nightmare"`
	Hell      Location `json:"hell"`
}

func NewLocations(r io.Reader) (*Locations, error) {
	var packedLocations [locationsCount]byte
	if err := binary.Read(r, binaryEndian, &packedLocations); err != nil {
		return nil, err
	}
	return &Locations{
		*UnpackLocation(packedLocations[0]),
		*UnpackLocation(packedLocations[1]),
		*UnpackLocation(packedLocations[2]),
	}, nil
}

//GetPacked returns packed locations to 3 bytes
func (l *Locations) GetPacked() [locationsCount]byte {
	return [3]byte{l.Normal.GetPacked(), l.Nightmare.GetPacked(), l.Hell.GetPacked()}
}
