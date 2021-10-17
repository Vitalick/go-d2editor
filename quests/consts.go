package quests

import (
	"github.com/vitalick/go-d2editor/consts"
)

var (
	binaryEndian = consts.BinaryEndian
)

//ActQuest types for indexing Quest in Act
type ActQuest byte

//Act 1 quests
const (
	Act1Introduction ActQuest = iota
	Act1DenOfEvil
	Act1SistersBurialGrounds
	Act1ToolsOfTheTrade
	Act1TheSearchForCain
	Act1TheForgottenTower
	Act1SistersToTheSlaughter
	Act1Completion
)

//Act 2 quests
const (
	Act2Introduction ActQuest = iota
	Act2RadamentsLair
	Act2TheHoradricStaff
	Act2TaintedSun
	Act2ArcaneSanctuary
	Act2TheSummoner
	Act2TheSevenTombs
	Act2Completion
)

//Act 3 quests
const (
	Act3Introduction ActQuest = iota
	Act3LamEsensTome
	Act3KhalimsWill
	Act3BladeOfTheOldReligion
	Act3TheGoldenBird
	Act3TheBlackenedTemple
	Act3TheGuardian
	Act3Completion
)

//Act 4 quests
const (
	Act4Introduction ActQuest = iota
	Act4TheFallenAngel
	Act4TerrorsEnd
	Act4Hellforge
	Act4Completion

	// Act4Extra1 Act4Extra2 Act4Extra3 3 shorts at the end of Act4 completion. presumably for extra quests never used.
	Act4Extra1
	Act4Extra2
	Act4Extra3
)

//Act 5 quests
const (
	Act5Introduction ActQuest = iota

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

var actQuestsMap = [][]string{
	{
		"Introduction",
		"Den Of Evil",
		"Sisters Burial Grounds",
		"Tools Of The Trade",
		"The Search For Cain",
		"The Forgotten Tower",
		"Sisters To The Slaughter",
		"Completion",
	},
	{
		"Introduction",
		"Radaments Lair",
		"The Horadric Staff",
		"Tainted Sun",
		"Arcane Sanctuary",
		"The Summoner",
		"The Seven Tombs",
		"Completion",
	},
	{
		"Introduction",
		"Lam Esens Tome",
		"Khalims Will",
		"Blade Of The Old Religion",
		"The Golden Bird",
		"The Blackened Temple",
		"The Guardian",
		"Completion",
	},
	{
		"Introduction",
		"The Fallen Angel",
		"Terrors End",
		"Hellforge",
		"Completion",
		"Extra 1",
		"Extra 2",
		"Extra 3",
	},
	{
		"Introduction",
		"Extra 1",
		"Extra 2",
		"Siege On Harrogath",
		"Rescue On Mount Arreat",
		"Prison Of Ice",
		"Betrayal Of Harrogath",
		"Rite Of Passage",
		"Eve Of Destruction",
		"Completion",
		"Extra 3",
		"Extra 4",
		"Extra 5",
		"Extra 6",
		"Extra 7",
		"Extra 8",
	},
}

var actLengths = []int{
	len(actQuestsMap[consts.Act1]),
	len(actQuestsMap[consts.Act2]),
	len(actQuestsMap[consts.Act3]),
	len(actQuestsMap[consts.Act4]),
	len(actQuestsMap[consts.Act5]),
}

//QuestFlag ...
type QuestFlag byte

//Types of QuestFlag
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

var questFlagsMap = []string{
	"Reward Granted",
	"Reward Pending",
	"Started",
	"Left Town",
	"Enter Area",
	"Custom 1",
	"Custom 2",
	"Custom 3",
	"Custom 4",
	"Custom 5",
	"Custom 6",
	"Custom 7",
	"Quest Log",
	"Primary Goal Achieved",
	"Completed Now",
	"Completed Before",
}

func (qf QuestFlag) String() string {
	return questFlagsMap[qf]
}

var questFlagCount = len(questFlagsMap)
