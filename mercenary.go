package d2s

type Mercenary struct {
	IsDead     uint16 `json:"is_dead"`
	Id         uint32 `json:"id"`
	NameId     uint16 `json:"name_id"`
	TypeId     uint16 `json:"type_id"`
	Experience uint32 `json:"experience"`
}
