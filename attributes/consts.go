package attributes

import (
	"github.com/vitalick/go-d2editor/consts"
)

var (
	binaryEndian = consts.BinaryEndian
)

const attributeMaxSize uint = 32

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
