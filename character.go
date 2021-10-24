package d2editor

import (
	"bytes"
	"encoding/binary"
	"github.com/vitalick/go-d2editor/attributes"
	"github.com/vitalick/go-d2editor/consts"
	"github.com/vitalick/go-d2editor/npcdialogs"
	"github.com/vitalick/go-d2editor/quests"
	"github.com/vitalick/go-d2editor/waypoints"
	"io"
	"time"
)

//Character ...
type Character struct {
	Header         *Header                `json:"header"`
	ActiveWeapon   uint32                 `json:"active_weapon"`
	Name           string                 `json:"name"`
	Status         *Status                `json:"status"`
	Progression    byte                   `json:"-"`
	Unk0x0026      [2]byte                `json:"-"`
	Class          class                  `json:"class"`
	Unk0x0029      [2]byte                `json:"-"`
	Level          byte                   `json:"level"`
	Created        uint32                 `json:"created"`
	LastPlayed     uint32                 `json:"last_played"`
	Unk0x0034      [4]byte                `json:"-"`
	HotkeySkills   [16]Skill              `json:"hotkey_skills"`
	LeftSkill      Skill                  `json:"left_skill"`
	RightSkill     Skill                  `json:"right_skill"`
	LeftSwapSkill  Skill                  `json:"left_swap_skill"`
	RightSwapSkill Skill                  `json:"right_swap_skill"`
	Appearances    Appearances            `json:"appearances"`
	Locations      *Locations             `json:"locations"`
	MapID          uint32                 `json:"map_id"`
	Unk0x00af      [2]byte                `json:"-"`
	Mercenary      Mercenary              `json:"mercenary"`
	RealmData      [144]byte              `json:"-"`
	Quests         *quests.Quests         `json:"quests"`
	Waypoints      *waypoints.Waypoints   `json:"waypoints"`
	UnkUnk1        byte                   `json:"-"`
	NPCDialogs     *npcdialogs.NPCDialogs `json:"npc_dialogs"`
	Attributes     *attributes.Attributes `json:"attributes"`
}

type inputStruct struct {
	data interface{}
	f    func(r io.Reader, c *Character) error
}

//NewEmptyCharacter ...
func NewEmptyCharacter(version uint) (*Character, error) {
	c := &Character{}
	c.Header = NewEmptyHeader(version)
	c.Status = &Status{}
	c.Locations = &Locations{}

	c.Level = 1
	timeNow := uint32(time.Now().Unix())
	c.Created = timeNow
	c.LastPlayed = timeNow

	emptySkill := NewEmptySkill()

	for s := range c.HotkeySkills {
		c.HotkeySkills[s] = *emptySkill
	}
	c.LeftSkill = *emptySkill
	c.RightSkill = *emptySkill
	c.LeftSwapSkill = *emptySkill
	c.RightSwapSkill = *emptySkill

	c.Appearances = *NewEmptyAppearances()
	q, err := quests.NewEmptyQuests()
	if err != nil {
		return nil, err
	}
	c.Quests = q

	w, err := waypoints.NewEmptyWaypoints()
	if err != nil {
		return nil, err
	}
	c.Waypoints = w
	c.NPCDialogs = npcdialogs.NewEmptyNPCDialogs()
	c.Attributes = attributes.NewEmptyAttributes()

	c.Unk0x0026 = defaultUnk0x0026
	c.Unk0x0029 = defaultUnk0x0029
	c.Unk0x0034 = defaultUnk0x0034
	c.Unk0x00af = defaultUnk0x00af
	c.RealmData = defaultRealmData

	return c, nil
}

//NewCharacter ...
func NewCharacter(r io.Reader) (*Character, error) {
	c := &Character{}
	var err error
	var charName [nameSize]byte

	var inArr []inputStruct

	inArr = append(inArr, inputStruct{
		nil,
		func(r io.Reader, c *Character) error {
			header, er := NewHeader(r)
			if er != nil {
				return er
			}
			c.Header = header
			return nil
		},
	})

	inArr = append(inArr, inputStruct{&c.ActiveWeapon, nil})
	inArr = append(inArr, inputStruct{&charName, nil})

	inArr = append(inArr, inputStruct{
		nil,
		func(r io.Reader, c *Character) error {
			status, er := NewStatus(r)
			if er != nil {
				return er
			}
			c.Status = status
			return nil
		},
	})

	inArr = append(inArr, inputStruct{&c.Progression, nil})
	inArr = append(inArr, inputStruct{&c.Unk0x0026, nil})
	inArr = append(inArr, inputStruct{&c.Class, nil})
	inArr = append(inArr, inputStruct{&c.Unk0x0029, nil})
	inArr = append(inArr, inputStruct{&c.Level, nil})
	inArr = append(inArr, inputStruct{&c.Created, nil})
	inArr = append(inArr, inputStruct{&c.LastPlayed, nil})
	inArr = append(inArr, inputStruct{&c.Unk0x0034, nil})
	inArr = append(inArr, inputStruct{&c.HotkeySkills, nil})
	inArr = append(inArr, inputStruct{&c.LeftSkill, nil})
	inArr = append(inArr, inputStruct{&c.RightSkill, nil})
	inArr = append(inArr, inputStruct{&c.LeftSwapSkill, nil})
	inArr = append(inArr, inputStruct{&c.RightSwapSkill, nil})
	inArr = append(inArr, inputStruct{&c.Appearances, nil})

	inArr = append(inArr, inputStruct{
		nil,
		func(r io.Reader, c *Character) error {
			loc, er := NewLocations(r)
			if er != nil {
				return er
			}
			c.Locations = loc
			return nil
		},
	})

	inArr = append(inArr, inputStruct{&c.MapID, nil})
	inArr = append(inArr, inputStruct{&c.Unk0x00af, nil})
	inArr = append(inArr, inputStruct{&c.Mercenary, nil})
	inArr = append(inArr, inputStruct{&c.RealmData, nil})

	inArr = append(inArr, inputStruct{
		nil,
		func(r io.Reader, c *Character) error {
			q, er := quests.NewQuests(r)
			if er != nil {
				return er
			}
			c.Quests = q
			return nil
		},
	})

	inArr = append(inArr, inputStruct{
		nil,
		func(r io.Reader, c *Character) error {
			w, er := waypoints.NewWaypoints(r)
			if er != nil {
				return er
			}
			c.Waypoints = w
			return nil
		},
	})
	inArr = append(inArr, inputStruct{&c.UnkUnk1, nil})

	inArr = append(inArr, inputStruct{
		nil,
		func(r io.Reader, c *Character) error {
			d, er := npcdialogs.NewNPCDialogs(r)
			if er != nil {
				return er
			}
			c.NPCDialogs = d
			return nil
		},
	})

	inArr = append(inArr, inputStruct{
		nil,
		func(r io.Reader, c *Character) error {
			a, er := attributes.NewAttributes(r)
			if er != nil {
				return er
			}
			c.Attributes = a
			return nil
		},
	})

	for _, inData := range inArr {
		if inData.f != nil {
			err = inData.f(r, c)
			if err != nil {
				return nil, err
			}
			continue
		}
		if inData.data != nil {
			err = binary.Read(r, consts.BinaryEndian, inData.data)
			if err != nil {
				return nil, err
			}
		}
	}

	c.Name = string(bytes.Trim(charName[:], "\x00"))

	c.Unk0x0026 = defaultUnk0x0026
	c.Unk0x0029 = defaultUnk0x0029
	c.Unk0x0034 = defaultUnk0x0034
	c.Unk0x00af = defaultUnk0x00af
	c.RealmData = defaultRealmData

	return c, nil
}

func (c *Character) GetName() string {
	if len(c.Name) > nameSize {
		return c.Name[:nameSize]
	}
	return c.Name
}

func (c *Character) getNameBytes() [nameSize]byte {
	var charName [nameSize]byte
	copy(charName[:], c.GetName()[:])
	return charName
}

//ToWriter write not prepared for export byte struct to io.Writer
func (c *Character) ToWriter(w io.Writer) error {
	c.Attributes.Level = uint64(c.Level)
	if err := c.Attributes.CheckMaxGold(); err != nil {
		return err
	}
	ww := &writerWrapper{w: w}
	var values [28]interface{}

	type packedChan struct {
		result []byte
		err    error
	}

	locationsCh := make(chan packedChan)
	questsCh := make(chan packedChan)
	waypointsCh := make(chan packedChan)
	dialogsCh := make(chan packedChan)
	attributesCh := make(chan packedChan)

	getPackedChan := func(f func() ([]byte, error), ch chan packedChan) {
		b, err := f()
		if err != nil {
			ch <- packedChan{nil, err}
			return
		}
		ch <- packedChan{b, nil}
	}
	go getPackedChan(func() ([]byte, error) { return c.Locations.GetPacked(), nil }, locationsCh)
	go getPackedChan(c.Quests.GetPacked, questsCh)
	go getPackedChan(c.Waypoints.GetPacked, waypointsCh)
	go getPackedChan(c.NPCDialogs.GetPacked, dialogsCh)
	go getPackedChan(c.Attributes.GetPacked, attributesCh)

	values[0] = *c.Header
	values[1] = c.ActiveWeapon
	values[2] = c.getNameBytes()
	values[3] = c.Status.GetFlags()
	values[4] = c.Progression
	values[5] = c.Unk0x0026
	values[6] = c.Class
	values[7] = c.Unk0x0029
	values[8] = c.Level
	values[9] = c.Created
	values[10] = c.LastPlayed
	values[11] = c.Unk0x0034
	values[12] = c.HotkeySkills
	values[13] = c.LeftSkill
	values[14] = c.RightSkill
	values[15] = c.LeftSwapSkill
	values[16] = c.RightSwapSkill
	values[17] = c.Appearances
	values[19] = c.MapID
	values[20] = c.Unk0x00af
	values[21] = c.Mercenary
	values[22] = c.RealmData
	values[25] = c.UnkUnk1

	for range make([]struct{}, 5) {
		select {
		case locationsB := <-locationsCh:
			values[18] = locationsB.result
		case questsB := <-questsCh:
			if questsB.err != nil {
				return questsB.err
			}
			values[23] = questsB.result
		case waypointsB := <-waypointsCh:
			if waypointsB.err != nil {
				return waypointsB.err
			}
			values[24] = waypointsB.result
		case dialogsB := <-dialogsCh:
			if dialogsB.err != nil {
				return dialogsB.err
			}
			values[26] = dialogsB.result
		case attributesB := <-attributesCh:
			if attributesB.err != nil {
				return attributesB.err
			}
			values[27] = attributesB.result
		}
	}

	for _, val := range values {
		if err := binary.Write(ww, consts.BinaryEndian, val); err != nil {
			return err
		}
	}
	if _, err := ww.EndWrite(); err != nil {
		return err
	}
	return nil
}

//GetBytes return not prepared for export []byte
func (c *Character) GetBytes() ([]byte, error) {
	var buf bytes.Buffer
	if err := c.ToWriter(&buf); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
