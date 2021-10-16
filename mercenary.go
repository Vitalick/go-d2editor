package d2editor

//Mercenary ...
type Mercenary struct {
	IsDead     uint16 `json:"is_dead"`
	ID         uint32 `json:"id"`
	NameID     uint16 `json:"name_id"`
	TypeID     uint16 `json:"type_id"`
	Experience uint32 `json:"experience"`
}
