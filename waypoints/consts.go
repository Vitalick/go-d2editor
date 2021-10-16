package waypoints

import (
	"github.com/vitalick/d2s/consts"
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

var waypointsMap = map[ActWaypoint]struct {
	name string
	act  consts.ActID
}{
	Act1RogueEncampement:       {"Rogue Encampement", consts.Act1},
	Act1ColdPlains:             {"Cold Plains", consts.Act1},
	Act1StonyField:             {"Stony Field", consts.Act1},
	Act1DarkWoods:              {"Dark Woods", consts.Act1},
	Act1BlackMarsh:             {"Black Marsh", consts.Act1},
	Act1OuterCloister:          {"Outer Cloister", consts.Act1},
	Act1JailLvl1:               {"Jail Lvl 1", consts.Act1},
	Act1InnerCloister:          {"Inner Cloister", consts.Act1},
	Act1CatacombsLvl2:          {"Catacombs Lvl 2", consts.Act1},
	Act2LutGholein:             {"Lut Gholein", consts.Act2},
	Act2SewersLvl2:             {"Sewers Lvl 2", consts.Act2},
	Act2DryHills:               {"Dry Hills", consts.Act2},
	Act2HallsOfTheDeadLvl2:     {"Halls Of The Dead Lvl 2", consts.Act2},
	Act2FarOasis:               {"Far Oasis", consts.Act2},
	Act2LostCity:               {"Lost City", consts.Act2},
	Act2PalaceCellarLvl1:       {"Palace Cellar Lvl 1", consts.Act2},
	Act2ArcaneSanctuary:        {"Arcane Sanctuary", consts.Act2},
	Act2CanyonOfTheMagi:        {"Canyon Of The Magi", consts.Act2},
	Act3KurastDocks:            {"Kurast Docks", consts.Act3},
	Act3SpiderForest:           {"Spider Forest", consts.Act3},
	Act3GreatMarsh:             {"Great Marsh", consts.Act3},
	Act3FlayerJungle:           {"Flayer Jungle", consts.Act3},
	Act3LowerKurast:            {"Lower Kurast", consts.Act3},
	Act3KurastBazaar:           {"Kurast Bazaar", consts.Act3},
	Act3UpperKurast:            {"Upper Kurast", consts.Act3},
	Act3Travincal:              {"Travincal", consts.Act3},
	Act3DuranceOfHateLvl2:      {"Durance Of Hate Lvl 2", consts.Act3},
	Act4ThePandemoniumFortress: {"The Pandemonium Fortress", consts.Act4},
	Act4CityOfTheDamned:        {"City Of The Damned", consts.Act4},
	Act4RiverOfFlame:           {"River Of Flame", consts.Act4},
	Act5Harrogath:              {"Harrogath", consts.Act5},
	Act5FrigidHighlands:        {"Frigid Highlands", consts.Act5},
	Act5ArreatPlateau:          {"Arreat Plateau", consts.Act5},
	Act5CrystallinePassage:     {"Crystalline Passage", consts.Act5},
	Act5HallsOfPain:            {"Halls Of Pain", consts.Act5},
	Act5GlacialTrail:           {"Glacial Trail", consts.Act5},
	Act5FrozenTundra:           {"Frozen Tundra", consts.Act5},
	Act5TheAncientsWay:         {"The Ancients Way", consts.Act5},
	Act5WorldstoneKeepLvl2:     {"Worldstone Keep Lvl 2", consts.Act5},
}

var actWaypointsCount = len(waypointsMap)

var actWaypointsMap = map[consts.ActID][]ActWaypoint{
	consts.Act1: {
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
	consts.Act2: {
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
	consts.Act3: {
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
	consts.Act4: {
		Act4ThePandemoniumFortress,
		Act4CityOfTheDamned,
		Act4RiverOfFlame,
	},
	consts.Act5: {
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
