package d2editor

//Appearances ...
type Appearances struct {
	Head      Appearance `json:"head,omitempty"`
	Torso     Appearance `json:"torso,omitempty"`
	Legs      Appearance `json:"legs,omitempty"`
	RightArm  Appearance `json:"right_arm,omitempty"`
	LeftArm   Appearance `json:"left_arm,omitempty"`
	RightHand Appearance `json:"right_hand,omitempty"`
	LeftHand  Appearance `json:"left_hand,omitempty"`
	Shield    Appearance `json:"shield,omitempty"`
	Special1  Appearance `json:"special_1,omitempty"`
	Special2  Appearance `json:"special_2,omitempty"`
	Special3  Appearance `json:"special_3,omitempty"`
	Special4  Appearance `json:"special_4,omitempty"`
	Special5  Appearance `json:"special_5,omitempty"`
	Special6  Appearance `json:"special_6,omitempty"`
	Special7  Appearance `json:"special_7,omitempty"`
	Special8  Appearance `json:"special_8,omitempty"`
}

func NewEmptyAppearances() *Appearances {
	emptyAppearance := NewEmptyAppearance()
	return &Appearances{
		Head:      *emptyAppearance,
		Torso:     *emptyAppearance,
		Legs:      *emptyAppearance,
		RightArm:  *emptyAppearance,
		LeftArm:   *emptyAppearance,
		RightHand: *emptyAppearance,
		LeftHand:  *emptyAppearance,
		Shield:    *emptyAppearance,
		Special1:  *emptyAppearance,
		Special2:  *emptyAppearance,
		Special3:  *emptyAppearance,
		Special4:  *emptyAppearance,
		Special5:  *emptyAppearance,
		Special6:  *emptyAppearance,
		Special7:  *emptyAppearance,
		Special8:  *emptyAppearance,
	}
}
