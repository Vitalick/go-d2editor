package attributes

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"github.com/vitalick/bitslice"
	"io"
)

const (
	defaultHeaderString   = "gf"
	headerLength          = 2
	indicatorsLengthBytes = 2
	indicatorsLength      = indicatorsLengthBytes * 8
)

var (
	defaultHeader   = [headerLength]byte{}
	wrongHeader     = errors.New("wrong attributes header")
	wrongIndicators = errors.New("wrong indicators length")
)

func init() {
	copy(defaultHeader[:], defaultHeaderString[:])
}

//Attributes ...
type Attributes struct {
	header        [headerLength]byte
	indicators    []bool
	Strength      int32        `json:"strength"`
	Energy        int32        `json:"energy"`
	Dexterity     int32        `json:"dexterity"`
	Vitality      int32        `json:"vitality"`
	StatPoints    int32        `json:"stat_points"`
	SkillChoices  int32        `json:"skill_choices"`
	Life          *CurrentBase `json:"life"`
	Mana          *CurrentBase `json:"mana"`
	Stamina       *CurrentBase `json:"stamina"`
	Level         int32        `json:"-"`
	Experience    int32        `json:"experience"`
	GoldInventory int32        `json:"gold_inventory"`
	GoldStash     int32        `json:"gold_stash"`
}

//NewEmptyAttributes returns empty Attributes
func NewEmptyAttributes() *Attributes {
	return &Attributes{
		header:     defaultHeader,
		indicators: make([]bool, indicatorsLength),
		Life:       &CurrentBase{},
		Mana:       &CurrentBase{},
		Stamina:    &CurrentBase{},
	}
}

func (a *Attributes) getInt32(r io.Reader, i *int, p interface{}) error {
	pn := p.(*int32)
	var u int32
	nowIter := *i + 0
	*i++
	if !a.indicators[nowIter] && false {
		return nil
	}
	if err := binary.Read(r, binaryEndian, &u); err != nil {
		return err
	}
	*pn = u
	return nil
}

func (a *Attributes) getCurrentBase(r io.Reader, i *int, p interface{}) error {
	pn := p.(*CurrentBase)
	cb, err := NewCurrentBase(r, a.indicators, i)
	if err != nil {
		return err
	}
	*pn = *cb
	return nil
}

//NewAttributes returns Attributes from packed bytes
func NewAttributes(r io.Reader) (*Attributes, error) {
	a := NewEmptyAttributes()
	if err := binary.Read(r, binaryEndian, &a.header); err != nil {
		return nil, err
	}
	headerString := string(bytes.Trim(a.header[:], "\x00"))
	if headerString != defaultHeaderString {
		return nil, wrongHeader
	}
	bs, err := bitslice.NewBitSliceFromReader(r, binaryEndian, indicatorsLengthBytes)
	if err != nil {
		return nil, err
	}

	a.indicators = bs.Slice
	if len(a.indicators) < indicatorsLength {
		return nil, wrongIndicators
	}

	qs := []struct {
		p interface{}
		f func(io.Reader, *int, interface{}) error
	}{
		{p: &a.Strength, f: a.getInt32},
		{p: &a.Energy, f: a.getInt32},
		{p: &a.Dexterity, f: a.getInt32},
		{p: &a.Vitality, f: a.getInt32},
		{p: &a.StatPoints, f: a.getInt32},
		{p: &a.SkillChoices, f: a.getInt32},
		{p: a.Life, f: a.getCurrentBase},
		{p: a.Mana, f: a.getCurrentBase},
		{p: a.Stamina, f: a.getCurrentBase},
		{p: &a.Level, f: a.getInt32},
		{p: &a.Experience, f: a.getInt32},
		{p: &a.GoldInventory, f: a.getInt32},
		{p: &a.GoldStash, f: a.getInt32},
	}
	for i, qn := range qs {
		if err = qn.f(r, &i, qn.p); err != nil {
			return nil, err
		}
		fmt.Println(i)
	}
	return a, nil
}

func (a *Attributes) CalcMaxGoldInventory() int32 {
	return a.Level * 10000
}

func (a *Attributes) CalcMaxGoldStash() int32 {
	if a.Level <= 30 {
		return (a.Level/10 + 1) * 50000
	}
	return (a.Level/2 + 1) * 50000
}

func (a *Attributes) CheckMaxGoldInventory() bool {
	return a.GoldInventory <= a.CalcMaxGoldInventory()
}

func (a *Attributes) CheckMaxGoldStash() bool {
	return a.GoldStash <= a.CalcMaxGoldStash()
}

func (a *Attributes) CheckMaxLevel() bool {
	return a.Level < 100
}

func (a *Attributes) CheckMaxGold() error {
	if !a.CheckMaxLevel() {
		return errors.New("level is too high")
	}
	if !a.CheckMaxGoldInventory() {
		return errors.New("inventory gold is too high")
	}
	if !a.CheckMaxGoldStash() {
		return errors.New("stash gold is too high")
	}
	return nil
}

func (a *Attributes) getInt32Packed(w io.Writer, v interface{}, i *int) error {
	val := v.(*int32)
	nowIter := *i + 0
	*i++
	b := *val != 0
	a.indicators[nowIter] = b
	if !b {
		return nil
	}
	if err := binary.Write(w, binaryEndian, *val); err != nil {
		return err
	}
	return nil
}

func (a *Attributes) getCurrentBasePacked(w io.Writer, v interface{}, i *int) error {
	val := v.(*CurrentBase)
	nowIter := *i + 0
	*i += 2
	b := val.Current != 0
	a.indicators[nowIter] = b
	if b {
		bts, err := val.Current.GetPacked()
		if err != nil {
			return err
		}
		if err = binary.Write(w, binaryEndian, bts); err != nil {
			return err
		}
	}
	b = val.Base != 0
	a.indicators[nowIter+1] = b
	if b {
		bts, err := val.Current.GetPacked()
		if err != nil {
			return err
		}
		if err = binary.Write(w, binaryEndian, bts); err != nil {
			return err
		}
	}
	return nil
}

//GetPacked returns packed Attributes into []byte
func (a *Attributes) GetPacked() ([]byte, error) {
	var buf bytes.Buffer
	a.indicators = make([]bool, indicatorsLength)
	qs := []struct {
		p interface{}
		f func(io.Writer, interface{}, *int) error
	}{
		{p: &a.Strength, f: a.getInt32Packed},
		{p: &a.Energy, f: a.getInt32Packed},
		{p: &a.Dexterity, f: a.getInt32Packed},
		{p: &a.Vitality, f: a.getInt32Packed},
		{p: &a.StatPoints, f: a.getInt32Packed},
		{p: &a.SkillChoices, f: a.getInt32Packed},
		{p: a.Life, f: a.getCurrentBasePacked},
		{p: a.Mana, f: a.getCurrentBasePacked},
		{p: a.Stamina, f: a.getCurrentBasePacked},
		{p: &a.Level, f: a.getInt32Packed},
		{p: &a.Experience, f: a.getInt32Packed},
		{p: &a.GoldInventory, f: a.getInt32Packed},
		{p: &a.GoldStash, f: a.getInt32Packed},
	}
	for i, qn := range qs {
		if err := qn.f(&buf, qn.p, &i); err != nil {
			return nil, err
		}
	}
	bts := buf.Bytes()
	buf = bytes.Buffer{}
	if err := binary.Write(&buf, binaryEndian, a.header); err != nil {
		return nil, err
	}

	bs := bitslice.NewBitSliceFromBool(a.indicators, binaryEndian)
	if err := bs.ToBuffer(&buf); err != nil {
		return nil, err
	}
	if err := binary.Write(&buf, binaryEndian, bts); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
