package d2editor

import (
	"github.com/vitalick/go-d2editor/bitworker"
)

const locationsCount = 3

//Locations struct shown opened acts and difficulties information
type Locations struct {
	Normal    Location `json:"normal"`
	Nightmare Location `json:"nightmare"`
	Hell      Location `json:"hell"`
}

//NewLocations returns Locations from io.Reader
func NewLocations(br *bitworker.BitReader) (*Locations, error) {
	packedLocations, err := br.ReadNextBitsByteSlice(locationsCount)
	if err != nil {
		return nil, err
	}
	return &Locations{
		*UnpackLocation(packedLocations[0]),
		*UnpackLocation(packedLocations[1]),
		*UnpackLocation(packedLocations[2]),
	}, nil
}

//GetPacked returns packed locations to bytes
func (l *Locations) GetPacked() []byte {
	return []byte{l.Normal.GetPacked(), l.Nightmare.GetPacked(), l.Hell.GetPacked()}
}
