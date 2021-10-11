package d2s

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"github.com/vitalick/d2s/consts"
	"github.com/vitalick/d2s/quests"
	"github.com/vitalick/d2s/waypoints"
	"io"
)

type Character struct {
	Header         *Header              `json:"header"`
	ActiveWeapon   uint32               `json:"active_weapon"`
	Name           string               `json:"name"`
	Status         *Status              `json:"status"`
	Progression    byte                 `json:"-"`
	Unk0x0026      [2]byte              `json:"-"`
	ClassId        byte                 `json:"class_id"`
	Unk0x0029      [2]byte              `json:"-"`
	Level          byte                 `json:"level"`
	Created        uint32               `json:"created"`
	LastPlayed     uint32               `json:"last_played"`
	Unk0x0034      [4]byte              `json:"-"`
	HotkeySkills   [16]Skill            `json:"hotkey_skills"`
	LeftSkill      Skill                `json:"left_skill"`
	RightSkill     Skill                `json:"right_skill"`
	LeftSwapSkill  Skill                `json:"left_swap_skill"`
	RightSwapSkill Skill                `json:"right_swap_skill"`
	Appearances    Appearances          `json:"appearances"`
	Locations      *Locations           `json:"locations"`
	MapId          uint32               `json:"map_id"`
	Unk0x00af      [2]byte              `json:"-"`
	Mercenary      Mercenary            `json:"mercenary"`
	RealmData      [144]byte            `json:"-"`
	Quests         *quests.Quests       `json:"quests"`
	Waypoints      *waypoints.Waypoints `json:"waypoints"`
}

type inputStruct struct {
	data interface{}
	f    func(r io.Reader, c *Character) error
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
	inArr = append(inArr, inputStruct{&c.ClassId, nil})
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

	inArr = append(inArr, inputStruct{&c.MapId, nil})
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

	c.Unk0x0026 = [2]byte{0x0, 0x0}
	c.Unk0x0029 = [2]byte{0x10, 0x1e}
	c.Unk0x0034 = [4]byte{0xff, 0xff, 0xff, 0xff}
	c.Unk0x00af = [2]byte{0x0, 0x0}
	c.RealmData = [144]byte{}

	return c, nil
}

//Fix changes struct for export
func (c *Character) Fix() error {
	if err := c.Header.Fix(c); err != nil {
		return err
	}
	return nil
}

//ToWriter write not prepared for export byte struct to io.Writer
func (c *Character) ToWriter(w io.Writer) error {
	var values []interface{}
	values = append(values, *c.Header)
	values = append(values, c.ActiveWeapon)
	var charName [nameSize]byte
	if len(c.Name) > nameSize {
		c.Name = c.Name[:nameSize]
	}
	copy(charName[:], c.Name[:])
	values = append(values, charName)
	values = append(values, c.Status.GetFlags())
	values = append(values, c.Progression)
	values = append(values, c.Unk0x0026)
	values = append(values, c.ClassId)
	values = append(values, c.Unk0x0029)
	values = append(values, c.Level)
	values = append(values, c.Created)
	values = append(values, c.LastPlayed)
	values = append(values, c.Unk0x0034)
	values = append(values, c.HotkeySkills)
	values = append(values, c.LeftSkill)
	values = append(values, c.RightSkill)
	values = append(values, c.LeftSwapSkill)
	values = append(values, c.RightSwapSkill)
	values = append(values, c.Appearances)
	values = append(values, c.Locations.GetPacked())
	values = append(values, c.MapId)
	values = append(values, c.Unk0x00af)
	values = append(values, c.Mercenary)
	values = append(values, c.RealmData)
	packedQuests, err := c.Quests.GetPacked()
	if err != nil {
		return err
	}
	values = append(values, packedQuests)
	packedWaypoints, err := c.Waypoints.GetPacked()
	if err != nil {
		return err
	}
	values = append(values, packedWaypoints)

	for _, val := range values {
		if err := binary.Write(w, consts.BinaryEndian, val); err != nil {
			return err
		}
	}
	return nil
}

//ToWriterCorrect write prepared for export byte struct to io.Writer
func (c *Character) ToWriterCorrect(w io.Writer) error {
	if err := c.Fix(); err != nil {
		return err
	}
	if err := c.ToWriter(w); err != nil {
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

//GetCorrectBytes return prepared for export []byte
func (c *Character) GetCorrectBytes() ([]byte, error) {
	var buf bytes.Buffer
	bw := bufio.NewWriter(&buf)
	if err := c.ToWriterCorrect(bw); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
