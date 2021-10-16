package d2s

import "math"

const emptyGraphic = math.MaxUint8
const emptyTint = emptyGraphic

//Appearance ...
type Appearance struct {
	Graphic byte `json:"graphic"`
	Tint    byte `json:"tint"`
}

func NewEmptyAppearance() *Appearance {
	a := &Appearance{}
	a.Clear()
	return a
}

func (a *Appearance) ClearGraphic() {
	a.Graphic = emptyGraphic
}

func (a *Appearance) ClearTint() {
	a.Tint = emptyTint
}

func (a *Appearance) Clear() {
	a.ClearGraphic()
	a.ClearTint()
}

func (a *Appearance) SetGraphic(val byte) {
	a.Graphic = val
}

func (a *Appearance) SetTint(val byte) {
	a.Tint = val
}

func (a *Appearance) Set(graphic, tint byte) {
	a.SetGraphic(graphic)
	a.SetTint(tint)
}
