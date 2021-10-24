package attributes

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
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
	Strength          uint64 `json:"strength"`
	Energy            uint64 `json:"energy"`
	Dexterity         uint64 `json:"dexterity"`
	Vitality          uint64 `json:"vitality"`
	UnusedStats       uint64 `json:"unused_stats"`
	UnusedSkillPoints uint64 `json:"unused_skill_points"`
	CurrentHP         uint64 `json:"current_hp"`
	MaxHP             uint64 `json:"max_hp"`
	CurrentMana       uint64 `json:"current_mana"`
	MaxMana           uint64 `json:"max_mana"`
	CurrentStamina    uint64 `json:"current_stamina"`
	MaxStamina        uint64 `json:"max_stamina"`
	Level             uint64 `json:"-"`
	Experience        uint64 `json:"experience"`
	GoldInventory     uint64 `json:"gold_inventory"`
	GoldStash         uint64 `json:"gold_stash"`
}

//NewEmptyAttributes returns empty Attributes
func NewEmptyAttributes() *Attributes {
	return &Attributes{
		header: defaultHeader,
	}
}

//NewAttributes returns Attributes from packed bytes
func NewAttributes(r io.Reader) (*Attributes, error) {
	a := NewEmptyAttributes()
	if err := binary.Read(r, binaryEndian, &a.header); err != nil {
		return nil, err
	}
	headerString := string(a.header[:])
	if headerString != defaultHeaderString {
		fmt.Printf("%02x\n", a.header)
		return nil, wrongHeader
	}
	lastByte := byte(0x0)
	var bytesA []byte
	for {
		var nowByte byte
		if err := binary.Read(r, binaryEndian, &nowByte); err != nil {
			return nil, err
		}
		bytesA = append(bytesA, nowByte)
		if nowByte == 0x3f && lastByte&0xe0 == 0xe0 {
			break
		}
		lastByte = nowByte
	}
	bitArr := newBitArray(bytesA)
	sumLen := 0
	for {
		id, err := bitArr.GetFirst(9)
		sumLen += 9
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
		valNew, err := bitArr.GetFirst(attrLen)
		sumLen += int(attrLen)
		//fmt.Println(attr)
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
			a.CurrentHP = valNew / 256
		case maxHP:
			a.MaxHP = valNew / 256
		case currentMana:
			a.CurrentMana = valNew / 256
		case maxMana:
			a.MaxMana = valNew / 256
		case currentStamina:
			a.CurrentStamina = valNew / 256
		case maxStamina:
			a.MaxStamina = valNew / 256
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
	buf := &bytes.Buffer{}
	if err := binary.Write(buf, binaryEndian, a.header); err != nil {
		return nil, err
	}
	outS := ""
	for i := range make([]struct{}, 16) {
		attr := AttributeID(i)
		attrLen := int(attr.Size())
		nowS := ""
		formatS := fmt.Sprintf("%%0%db", attrLen)
		switch attr {
		case strength:
			nowS = fmt.Sprintf(formatS, a.Strength)
		case energy:
			nowS = fmt.Sprintf(formatS, a.Energy)
		case dexterity:
			nowS = fmt.Sprintf(formatS, a.Dexterity)
		case vitality:
			nowS = fmt.Sprintf(formatS, a.Vitality)
		case unusedStats:
			nowS = fmt.Sprintf(formatS, a.UnusedStats)
		case unusedSkills:
			nowS = fmt.Sprintf(formatS, a.UnusedSkillPoints)
		case currentHP:
			nowS = fmt.Sprintf(formatS, a.CurrentHP*256)
		case maxHP:
			nowS = fmt.Sprintf(formatS, a.MaxHP*256)
		case currentMana:
			nowS = fmt.Sprintf(formatS, a.CurrentMana*256)
		case maxMana:
			nowS = fmt.Sprintf(formatS, a.MaxMana*256)
		case currentStamina:
			nowS = fmt.Sprintf(formatS, a.CurrentStamina*256)
		case maxStamina:
			nowS = fmt.Sprintf(formatS, a.MaxStamina*256)
		case level:
			nowS = fmt.Sprintf(formatS, a.Level)
		case experience:
			nowS = fmt.Sprintf(formatS, a.Experience)
		case gold:
			nowS = fmt.Sprintf(formatS, a.GoldInventory)
		case stashedGold:
			nowS = fmt.Sprintf(formatS, a.GoldStash)
		default:
			break
		}
		nowS = nowS[len(nowS)-attrLen:]
		//fmt.Println("attrLen", attrLen)
		//fmt.Println("nowS", nowS)
		//fmt.Println()
		cmpS := ""
		for range nowS {
			cmpS += "0"
		}
		if nowS == cmpS {
			continue
		}
		idCounter := fmt.Sprintf("%09b", attr)
		//fmt.Println("attrLen", 9)
		//fmt.Println("nowS", idCounter)
		//fmt.Println()
		outS = nowS + idCounter + outS
	}
	outS = "00111111111" + outS
	fmt.Println(outS)
	fmt.Println("outS", len(outS))
	ba := bitArray{s: outS}
	bts, err := ba.GetBytes()
	if err != nil {
		return nil, err
	}
	if err = binary.Write(buf, binaryEndian, bts); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
