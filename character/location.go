package character

//Location is information about difficult and opened acts on it
type Location struct {
	Active bool `json:"active"`
	Act    byte `json:"act"`
}

//NewLocation ...
func NewLocation(active bool, act byte) *Location {
	return &Location{
		active, act,
	}
}

//UnpackLocation returns Location from packed byte
func UnpackLocation(locationInfo byte) *Location {
	return NewLocation(locationInfo>>7 == 1, (locationInfo&0x5)+1)
}

//GetPacked returns packed location to one byte
func (l *Location) GetPacked() byte {
	var packed byte
	if l.Active {
		packed |= 1 << 7
	}
	packed = packed | l.Act - 1
	return packed
}
