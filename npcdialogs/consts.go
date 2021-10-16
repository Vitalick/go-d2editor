package npcdialogs

import (
	"github.com/vitalick/go-d2editor/consts"
)

var (
	binaryEndian = consts.BinaryEndian
)

//NPCDialog types for indexing NPCDialog in Difficulty
type NPCDialog byte

//NPCDialog list
const (
	DialogWarrivAct2 NPCDialog = iota
	DialogUnk0x0001
	DialogCharsi
	DialogWarrivAct1
	DialogKashya
	DialogAkara
	DialogGheed
	DialogUnk0x0007
	DialogGreiz
	DialogJerhyn
	DialogMeshifAct2
	DialogGeglash
	DialogLysander
	DialogFara
	DialogDrogan
	DialogUnk0x000F
	DialogAlkor
	DialogHratli
	DialogAshera
	DialogUnk0x0013
	DialogUnk0x0014
	DialogCainAct3
	DialogUnk0x0016
	DialogElzix
	DialogMalah
	DialogAnya
	DialogUnk0x001A
	DialogNatalya
	DialogMeshifAct3
	DialogUnk0x001D
	DialogUnk0x001F
	DialogOrmus
	DialogUnk0x0021
	DialogUnk0x0022
	DialogUnk0x0023
	DialogUnk0x0024
	DialogUnk0x0025
	DialogCainAct5
	DialogQualkehk
	DialogNihlathak
	DialogUnk0x0029
)

var npcDialogMap = map[NPCDialog]string{
	DialogWarrivAct2: "Warriv Act 2",
	DialogUnk0x0001:  "Unknown 0x0001",
	DialogCharsi:     "Charsi",
	DialogWarrivAct1: "Warriv Act 1",
	DialogKashya:     "Kashya",
	DialogAkara:      "Akara",
	DialogGheed:      "Gheed",
	DialogUnk0x0007:  "Unknown 0x0007",
	DialogGreiz:      "Greiz",
	DialogJerhyn:     "Jerhyn",
	DialogMeshifAct2: "Meshif Act 2",
	DialogGeglash:    "Geglash",
	DialogLysander:   "Lysander",
	DialogFara:       "Fara",
	DialogDrogan:     "Drogan",
	DialogUnk0x000F:  "Unknown 0x000F",
	DialogAlkor:      "Alkor",
	DialogHratli:     "Hratli",
	DialogAshera:     "Ashera",
	DialogUnk0x0013:  "Unknown 0x0013",
	DialogUnk0x0014:  "Unknown 0x0014",
	DialogCainAct3:   "Cain Act 3",
	DialogUnk0x0016:  "Unknown 0x0016",
	DialogElzix:      "Elzix",
	DialogMalah:      "Malah",
	DialogAnya:       "Anya",
	DialogUnk0x001A:  "Unknown 0x001A",
	DialogNatalya:    "Natalya",
	DialogMeshifAct3: "Meshif Act 3",
	DialogUnk0x001D:  "Unknown 0x001D",
	DialogUnk0x001F:  "Unknown 0x001F",
	DialogOrmus:      "Ormus",
	DialogUnk0x0021:  "Unknown 0x0021",
	DialogUnk0x0022:  "Unknown 0x0022",
	DialogUnk0x0023:  "Unknown 0x0023",
	DialogUnk0x0024:  "Unknown 0x0024",
	DialogUnk0x0025:  "Unknown 0x0025",
	DialogCainAct5:   "Cain Act 5",
	DialogQualkehk:   "Qualkehk",
	DialogNihlathak:  "Nihlathak",
	DialogUnk0x0029:  "Unknown 0x0029",
}

var npcDialogsCount = len(npcDialogMap)

func (n NPCDialog) String() string {
	return npcDialogMap[n]
}

const (
	bitSliceSizeBytes = 0x30
	dialogTypeOffset  = 0x18 * 8
)
