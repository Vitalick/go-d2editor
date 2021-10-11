package character

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"io"
)

type Character struct {
	Header         *Header     `json:"header"`
	ActiveWeapon   uint32      `json:"active_weapon"`
	Name           string      `json:"name"`
	Status         *Status     `json:"status"`
	Progression    byte        `json:"-"`
	Unk0x0026      [2]byte     `json:"-"`
	ClassId        byte        `json:"class_id"`
	Unk0x0029      [2]byte     `json:"-"`
	Level          byte        `json:"level"`
	Created        uint32      `json:"created"`
	LastPlayed     uint32      `json:"last_played"`
	Unk0x0034      [4]byte     `json:"-"`
	HotkeySkills   [16]Skill   `json:"hotkey_skills"`
	LeftSkill      Skill       `json:"left_skill"`
	RightSkill     Skill       `json:"right_skill"`
	LeftSwapSkill  Skill       `json:"left_swap_skill"`
	RightSwapSkill Skill       `json:"right_swap_skill"`
	Appearances    Appearances `json:"appearances"`
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

	for _, inData := range inArr {
		if inData.f != nil {
			err = inData.f(r, c)
			if err != nil {
				return nil, err
			}
			continue
		}
		if inData.data != nil {
			err = binary.Read(r, binaryEndian, inData.data)
			if err != nil {
				return nil, err
			}
		}
	}

	c.Name = string(bytes.Trim(charName[:], "\x00"))

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
	for _, val := range values {
		if err := binary.Write(w, binaryEndian, val); err != nil {
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
