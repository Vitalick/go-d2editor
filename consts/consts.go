package consts

import "encoding/binary"

//ActId type for quests.Act and waypoints.Act and for indexing Act in quests.Difficulty and waypoints.Difficulty
type ActId byte

var (
	BinaryEndian = binary.LittleEndian
)

const (
	Act1 ActId = iota
	Act2
	Act3
	Act4
	Act5
)

var actsMap = map[ActId]string{
	Act1: "Act 1",
	Act2: "Act 2",
	Act3: "Act 3",
	Act4: "Act 4",
	Act5: "Act 5",
}

var ActsCount = len(actsMap)

func (a ActId) String() string {
	return actsMap[a]
}
