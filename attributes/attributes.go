package attributes

import (
	"errors"
	"fmt"
	"github.com/vitalick/go-d2editor/bitworker"
)

const (
	defaultHeaderString = "gf"
	headerLength        = 2
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
	header            [headerLength]byte
	Strength          uint64  `json:"strength"`
	Energy            uint64  `json:"energy"`
	Dexterity         uint64  `json:"dexterity"`
	Vitality          uint64  `json:"vitality"`
	UnusedStats       uint64  `json:"unused_stats"`
	UnusedSkillPoints uint64  `json:"unused_skill_points"`
	CurrentHP         float64 `json:"current_hp"`
	MaxHP             float64 `json:"max_hp"`
	CurrentMana       float64 `json:"current_mana"`
	MaxMana           float64 `json:"max_mana"`
	CurrentStamina    float64 `json:"current_stamina"`
	MaxStamina        float64 `json:"max_stamina"`
	Level             uint64  `json:"-"`
	Experience        uint64  `json:"experience"`
	GoldInventory     uint64  `json:"gold_inventory"`
	GoldStash         uint64  `json:"gold_stash"`
}

//NewEmptyAttributes returns empty Attributes
func NewEmptyAttributes() *Attributes {
	return &Attributes{
		header: defaultHeader,
	}
}

//NewAttributes returns Attributes from packed bytes
func NewAttributes(br *bitworker.BitReader) (*Attributes, error) {
	a := NewEmptyAttributes()
	if err := br.ReadNextBitsByteArray(a.header[:]); err != nil {
		return nil, err
	}
	headerString := string(a.header[:])
	if headerString != defaultHeaderString {
		fmt.Printf("%02x\n", a.header)
		return nil, wrongHeader
	}
	for {
		id, err := br.ReadNextBits(9)
		if err != nil {
			return nil, err
		}
		if id == 0x1ff {
			break
		}
		attr := AttributeID(id)
		attrLen := attr.Size()
		if attrLen == 0 {
			return nil, errors.New(fmt.Sprintf("unknown attribute id: %d %b", id, id))
		}
		valNew, err := br.ReadNextBits(uint64(attrLen))
		if err != nil {
			return nil, err
		}
		switch attr {
		case strength:
			a.Strength = valNew
		case energy:
			a.Energy = valNew
		case dexterity:
			a.Dexterity = valNew
		case vitality:
			a.Vitality = valNew
		case unusedStats:
			a.UnusedStats = valNew
		case unusedSkills:
			a.UnusedSkillPoints = valNew
		case currentHP:
			a.CurrentHP = float64(valNew) / 256
		case maxHP:
			a.MaxHP = float64(valNew) / 256
		case currentMana:
			a.CurrentMana = float64(valNew) / 256
		case maxMana:
			a.MaxMana = float64(valNew) / 256
		case currentStamina:
			a.CurrentStamina = float64(valNew) / 256
		case maxStamina:
			a.MaxStamina = float64(valNew) / 256
		case level:
			a.Level = valNew
		case experience:
			a.Experience = valNew
		case gold:
			a.GoldInventory = valNew
		case stashedGold:
			a.GoldStash = valNew
		}
	}
	if err := br.MoveToNextByte(); err != nil {
		return nil, err
	}
	return a, nil
}

func (a *Attributes) CalcMaxGoldInventory() uint64 {
	return a.Level * 10000
}

func (a *Attributes) CalcMaxGoldStash() uint64 {
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

//GetPacked returns packed Attributes into []byte
func (a *Attributes) GetPacked() ([]byte, error) {
	bw := bitworker.NewBitWriter(nil)

	if err := bw.WriteNextBitsByteSlice(a.header[:]); err != nil {
		return nil, err
	}
	for i := range make([]struct{}, 16) {
		attr := AttributeID(i)
		attrLen := int(attr.Size())
		var nowV uint64
		switch attr {
		case strength:
			nowV = a.Strength
		case energy:
			nowV = a.Energy
		case dexterity:
			nowV = a.Dexterity
		case vitality:
			nowV = a.Vitality
		case unusedStats:
			nowV = a.UnusedStats
		case unusedSkills:
			nowV = a.UnusedSkillPoints
		case currentHP:
			nowV = uint64(a.CurrentHP * 256)
		case maxHP:
			nowV = uint64(a.MaxHP * 256)
		case currentMana:
			nowV = uint64(a.CurrentMana * 256)
		case maxMana:
			nowV = uint64(a.MaxMana * 256)
		case currentStamina:
			nowV = uint64(a.CurrentStamina * 256)
		case maxStamina:
			nowV = uint64(a.MaxStamina * 256)
		case level:
			nowV = a.Level
		case experience:
			nowV = a.Experience
		case gold:
			nowV = a.GoldInventory
		case stashedGold:
			nowV = a.GoldStash
		default:
			break
		}
		if nowV != 0 {
			if err := bw.WriteNextBits(uint64(i), 9); err != nil {
				return nil, err
			}
			if err := bw.WriteNextBits(nowV, uint64(attrLen)); err != nil {
				return nil, err
			}
		}
	}
	if err := bw.WriteNextBits(0x1ff, 9); err != nil {
		return nil, err
	}
	return bw.Bytes, nil

}
