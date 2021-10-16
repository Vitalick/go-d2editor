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

var actQuestsMap = map[consts.ActID]map[ActQuest]string{
	consts.Act1: {
		Act1Introduction:          "Introduction",
		Act1DenOfEvil:             "Den Of Evil",
		Act1SistersBurialGrounds:  "Sisters Burial Grounds",
		Act1ToolsOfTheTrade:       "Tools Of The Trade",
		Act1TheSearchForCain:      "The Search For Cain",
		Act1TheForgottenTower:     "The Forgotten Tower",
		Act1SistersToTheSlaughter: "Sisters To The Slaughter",
		Act1Completion:            "Completion",
	},
	consts.Act2: {
		Act2Introduction:     "Introduction",
		Act2RadamentsLair:    "Radaments Lair",
		Act2TheHoradricStaff: "The Horadric Staff",
		Act2TaintedSun:       "Tainted Sun",
		Act2ArcaneSanctuary:  "Arcane Sanctuary",
		Act2TheSummoner:      "The Summoner",
		Act2TheSevenTombs:    "The Seven Tombs",
		Act2Completion:       "Completion",
	},
	consts.Act3: {
		Act3Introduction:          "Introduction",
		Act3LamEsensTome:          "Lam Esens Tome",
		Act3KhalimsWill:           "Khalims Will",
		Act3BladeOfTheOldReligion: "Blade Of The Old Religion",
		Act3TheGoldenBird:         "The Golden Bird",
		Act3TheBlackenedTemple:    "The Blackened Temple",
		Act3TheGuardian:           "The Guardian",
		Act3Completion:            "Completion",
	},
	consts.Act4: {
		Act4Introduction:   "Introduction",
		Act4TheFallenAngel: "The Fallen Angel",
		Act4TerrorsEnd:     "Terrors End",
		Act4Hellforge:      "Hellforge",
		Act4Completion:     "Completion",
		Act4Extra1:         "Extra 1",
		Act4Extra2:         "Extra 2",
		Act4Extra3:         "Extra 3",
	},
	consts.Act5: {
		Act5Introduction:        "Introduction",
		Act5Extra1:              "Extra 1",
		Act5Extra2:              "Extra 2",
		Act5SiegeOnHarrogath:    "Siege On Harrogath",
		Act5RescueOnMountArreat: "Rescue On Mount Arreat",
		Act5PrisonOfIce:         "Prison Of Ice",
		Act5BetrayalOfHarrogath: "Betrayal Of Harrogath",
		Act5RiteOfPassage:       "Rite Of Passage",
		Act5EveOfDestruction:    "Eve Of Destruction",
		Act5Completion:          "Completion",
		Act5Extra3:              "Extra 3",
		Act5Extra4:              "Extra 4",
		Act5Extra5:              "Extra 5",
		Act5Extra6:              "Extra 6",
		Act5Extra7:              "Extra 7",
		Act5Extra8:              "Extra 8",
	},
}

var actLengths = map[consts.ActID]int{
	consts.Act1: len(actQuestsMap[consts.Act1]),
	consts.Act2: len(actQuestsMap[consts.Act2]),
	consts.Act3: len(actQuestsMap[consts.Act3]),
	consts.Act4: len(actQuestsMap[consts.Act4]),
	consts.Act5: len(actQuestsMap[consts.Act5]),
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

var questFlagsMap = map[QuestFlag]string{
	FlagRewardGranted:       "Reward Granted",
	FlagRewardPending:       "Reward Pending",
	FlagStarted:             "Started",
	FlagLeftTown:            "Left Town",
	FlagEnterArea:           "Enter Area",
	FlagCustom1:             "Custom 1",
	FlagCustom2:             "Custom 2",
	FlagCustom3:             "Custom 3",
	FlagCustom4:             "Custom 4",
	FlagCustom5:             "Custom 5",
	FlagCustom6:             "Custom 6",
	FlagCustom7:             "Custom 7",
	FlagQuestLog:            "Quest Log",
	FlagPrimaryGoalAchieved: "Primary Goal Achieved",
	FlagCompletedNow:        "Completed Now",
	FlagCompletedBefore:     "Completed Before",
}

func (qf QuestFlag) String() string {
	return questFlagsMap[qf]
}

var questFlagCount = len(questFlagsMap)
