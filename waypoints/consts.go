package waypoints

import (
	"github.com/vitalick/d2s/consts"
)

var (
	binaryEndian = consts.BinaryEndian
)

//Act1 Act2 Act3 Act4 Act5 types for indexing Quest in Act
type Act1 byte
type Act2 byte
type Act3 byte
type Act4 byte
type Act5 byte
type QuestFlag byte

const (
	Act1RogueEncampement Act1 = iota
	Act1ColdPlains
	Act1StonyField
	Act1DarkWoods
	Act1BlackMarsh
	Act1OuterCloister
	Act1JailLvl1
	Act1InnerCloister
	Act1CatacombsLvl2
)
const act1Size = 9

const (
	Act2LutGholein Act2 = iota
	Act2SewersLvl2
	Act2DryHills
	Act2HallsOfTheDeadLvl2
	Act2FarOasis
	Act2LostCity
	Act2PalaceCellarLvl1
	Act2ArcaneSanctuary
	Act2CanyonOfTheMagi
)
const act2Size = 9

const (
	Act3KurastDocks Act3 = iota
	Act3SpiderForest
	Act3GreatMarsh
	Act3FlayerJungle
	Act3LowerKurast
	Act3KurastBazaar
	Act3UpperKurast
	Act3Travincal
	Act3DuranceOfHateLvl2
)
const act3Size = 9

const (
	Act4ThePandemoniumFortress Act4 = iota
	Act4CityOfTheDamned
	Act4RiverOfFlame
)
const act4Size = 3

const (
	Act5Harrogath Act5 = iota
	Act5FrigidHighlands
	Act5ArreatPlateau
	Act5CrystallinePassage
	Act5HallsOfPain
	Act5GlacialTrail
	Act5FrozenTundra
	Act5TheAncientsWay
	Act5WorldstoneKeepLvl2
)
const act5Size = 9

const (
	FlagRewardGranted QuestFlag = iota
	FlagRewardPending
	FlagStarted
	FlagLeftTown
	FlagEnterArea
	FlagCustom1
	FlagCustom2
	FlagCustom3
	FlagCustom4
	FlagCustom5
	FlagCustom6
	FlagCustom7
	FlagQuestLog
	FlagPrimaryGoalAchieved
	FlagCompletedNow
	FlagCompletedBefore
)
const actWaypointsMaxCount = 8 * 2

var actCountTable = [consts.ActsCount]byte{act1Size, act2Size, act3Size, act4Size, act5Size}

const difficultySize = act1Size + act2Size + act3Size + act4Size + act5Size
