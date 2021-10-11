package quests

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
	Act1Introduction Act1 = iota
	Act1DenOfEvil
	Act1SistersBurialGrounds
	Act1ToolsOfTheTrade
	Act1TheSearchForCain
	Act1TheForgottenTower
	Act1SistersToTheSlaughter
	Act1Completion
)
const act1Size = 8

const (
	Act2Introduction Act2 = iota
	Act2RadamentsLair
	Act2TheHoradricStaff
	Act2TaintedSun
	Act2ArcaneSanctuary
	Act2TheSummoner
	Act2TheSevenTombs
	Act2Completion
)
const act2Size = 8

const (
	Act3Introduction Act3 = iota
	Act3LamEsensTome
	Act3KhalimsWill
	Act3BladeOfTheOldReligion
	Act3TheGoldenBird
	Act3TheBlackenedTemple
	Act3TheGuardian
	Act3Completion
)
const act3Size = 8

const (
	Act4Introduction Act4 = iota
	Act4TheFallenAngel
	Act4TerrorsEnd
	Act4Hellforge
	Act4Completion

	// Act4Extra1 Act4Extra2 Act4Extra3 3 shorts at the end of Act4 completion. presumably for extra quests never used.
	Act4Extra1
	Act4Extra2
	Act4Extra3
)
const act4Size = 8

const (
	Act5Introduction Act5 = iota

	// Act5Extra1 Act5Extra2 2 shorts after Act5 introduction. presumably for extra quests never used.
	Act5Extra1
	Act5Extra2

	Act5SiegeOnHarrogath
	Act5RescueOnMountArreat
	Act5PrisonOfIce
	Act5BetrayalOfHarrogath
	Act5RiteOfPassage
	Act5EveOfDestruction
	Act5Completion

	// Act5Extra3 Act5Extra4 Act5Extra5 Act5Extra6 Act5Extra7 Act5Extra8 6 shorts
	//after Act5 completion. presumably for extra quests never used.
	Act5Extra3
	Act5Extra4
	Act5Extra5
	Act5Extra6
	Act5Extra7
	Act5Extra8
)
const act5Size = 16

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
const questFlagCount = 16

var actCountTable = [consts.ActsCount]byte{act1Size, act2Size, act3Size, act4Size, act5Size}

const difficultySize = act1Size + act2Size + act3Size + act4Size + act5Size
