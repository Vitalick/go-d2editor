package consts

import "encoding/binary"

//ActID type for quests.Act and waypoints.Act and for indexing Act in quests.Difficulty and waypoints.Difficulty
type ActID byte

var (
	//BinaryEndian default binary endian for whole project
	BinaryEndian = binary.LittleEndian
)

//Act indexes
const (
	Act1 ActID = iota
	Act2
	Act3
	Act4
	Act5
)

var actsMap = map[ActID]string{
	Act1: "Act 1",
	Act2: "Act 2",
	Act3: "Act 3",
	Act4: "Act 4",
	Act5: "Act 5",
}

//ActsCount it is count of acts
var ActsCount = len(actsMap)

func (a ActID) String() string {
	return actsMap[a]
}
