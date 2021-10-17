package waypoints

import (
	"github.com/vitalick/go-d2editor/consts"
)

var (
	binaryEndian = consts.BinaryEndian
)

//ActWaypoint types for indexing Quest in Act
type ActWaypoint byte

//QuestFlag ...
type QuestFlag byte

//ActWaypoint list
const (
	Act1RogueEncampement ActWaypoint = iota
	Act1ColdPlains
	Act1StonyField
	Act1DarkWoods
	Act1BlackMarsh
	Act1OuterCloister
	Act1JailLvl1
	Act1InnerCloister
	Act1CatacombsLvl2

	Act2LutGholein
	Act2SewersLvl2
	Act2DryHills
	Act2HallsOfTheDeadLvl2
	Act2FarOasis
	Act2LostCity
	Act2PalaceCellarLvl1
	Act2ArcaneSanctuary
	Act2CanyonOfTheMagi

	Act3KurastDocks
	Act3SpiderForest
	Act3GreatMarsh
	Act3FlayerJungle
	Act3LowerKurast
	Act3KurastBazaar
	Act3UpperKurast
	Act3Travincal
	Act3DuranceOfHateLvl2

	Act4ThePandemoniumFortress
	Act4CityOfTheDamned
	Act4RiverOfFlame

	Act5Harrogath
	Act5FrigidHighlands
	Act5ArreatPlateau
	Act5CrystallinePassage
	Act5HallsOfPain
	Act5GlacialTrail
	Act5FrozenTundra
	Act5TheAncientsWay
	Act5WorldstoneKeepLvl2
)

var waypointsMap = []struct {
	name string
	act  consts.ActID
}{
	{"Rogue Encampement", consts.Act1},
	{"Cold Plains", consts.Act1},
	{"Stony Field", consts.Act1},
	{"Dark Woods", consts.Act1},
	{"Black Marsh", consts.Act1},
	{"Outer Cloister", consts.Act1},
	{"Jail Lvl 1", consts.Act1},
	{"Inner Cloister", consts.Act1},
	{"Catacombs Lvl 2", consts.Act1},
	{"Lut Gholein", consts.Act2},
	{"Sewers Lvl 2", consts.Act2},
	{"Dry Hills", consts.Act2},
	{"Halls Of The Dead Lvl 2", consts.Act2},
	{"Far Oasis", consts.Act2},
	{"Lost City", consts.Act2},
	{"Palace Cellar Lvl 1", consts.Act2},
	{"Arcane Sanctuary", consts.Act2},
	{"Canyon Of The Magi", consts.Act2},
	{"Kurast Docks", consts.Act3},
	{"Spider Forest", consts.Act3},
	{"Great Marsh", consts.Act3},
	{"Flayer Jungle", consts.Act3},
	{"Lower Kurast", consts.Act3},
	{"Kurast Bazaar", consts.Act3},
	{"Upper Kurast", consts.Act3},
	{"Travincal", consts.Act3},
	{"Durance Of Hate Lvl 2", consts.Act3},
	{"The Pandemonium Fortress", consts.Act4},
	{"City Of The Damned", consts.Act4},
	{"River Of Flame", consts.Act4},
	{"Harrogath", consts.Act5},
	{"Frigid Highlands", consts.Act5},
	{"Arreat Plateau", consts.Act5},
	{"Crystalline Passage", consts.Act5},
	{"Halls Of Pain", consts.Act5},
	{"Glacial Trail", consts.Act5},
	{"Frozen Tundra", consts.Act5},
	{"The Ancients Way", consts.Act5},
	{"Worldstone Keep Lvl 2", consts.Act5},
}

var actWaypointsCount = len(waypointsMap)

var actWaypointsMap = [][]ActWaypoint{
	{
		Act1RogueEncampement,
		Act1ColdPlains,
		Act1StonyField,
		Act1DarkWoods,
		Act1BlackMarsh,
		Act1OuterCloister,
		Act1JailLvl1,
		Act1InnerCloister,
		Act1CatacombsLvl2,
	},
	{
		Act2LutGholein,
		Act2SewersLvl2,
		Act2DryHills,
		Act2HallsOfTheDeadLvl2,
		Act2FarOasis,
		Act2LostCity,
		Act2PalaceCellarLvl1,
		Act2ArcaneSanctuary,
		Act2CanyonOfTheMagi,
	},
	{
		Act3KurastDocks,
		Act3SpiderForest,
		Act3GreatMarsh,
		Act3FlayerJungle,
		Act3LowerKurast,
		Act3KurastBazaar,
		Act3UpperKurast,
		Act3Travincal,
		Act3DuranceOfHateLvl2,
	},
	{
		Act4ThePandemoniumFortress,
		Act4CityOfTheDamned,
		Act4RiverOfFlame,
	},
	{
		Act5Harrogath,
		Act5FrigidHighlands,
		Act5ArreatPlateau,
		Act5CrystallinePassage,
		Act5HallsOfPain,
		Act5GlacialTrail,
		Act5FrozenTundra,
		Act5TheAncientsWay,
		Act5WorldstoneKeepLvl2,
	},
}

func (aw ActWaypoint) String() string {
	return waypointsMap[aw].name
}

//ActID returns consts.ActID
func (aw ActWaypoint) ActID() consts.ActID {
	return waypointsMap[aw].act
}
