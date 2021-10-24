package attributes

import (
	"github.com/vitalick/go-d2editor/consts"
)

var (
	binaryEndian = consts.BinaryEndian
)

type AttributeID uint16

const baseBitMask = 0xffffffffffffffff

const (
	strength AttributeID = iota
	energy
	dexterity
	vitality
	unusedStats
	unusedSkills
	currentHP
	maxHP
	currentMana
	maxMana
	currentStamina
	maxStamina
	level
	experience
	gold
	stashedGold
)

var attributesBitsAmount = []uint{
	10,
	10,
	10,
	10,
	10,
	8,
	21,
	21,
	21,
	21,
	21,
	21,
	7,
	32,
	25,
	25,
}

var attributesCount = len(attributesBitsAmount)

func (a AttributeID) Size() uint {
	if int(a) > attributesCount-1 {
		return 0
	}
	return attributesBitsAmount[a]
}
