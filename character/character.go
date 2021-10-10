package character

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"io"
)

type Character struct {
	Header       *Header
	ActiveWeapon uint32
	Name         string
	Status       *Status
	Progression  byte
	Unk0x0026    [2]byte `json:"-"`
	ClassId      byte
	Unk0x0029    [2]byte `json:"-"`
	Level        byte
	Created      uint32
	LastPlayed   uint32
	Unk0x0034    [4]byte `json:"-"`
}

type inputStruct struct {
	data interface{}
	f    func(r io.Reader, c *Character) error
}

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

func (c *Character) Fix() error {
	if err := c.Header.Fix(c); err != nil {
		return err
	}
	return nil
}

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

func (c *Character) ToWriterCorrect(w io.Writer) error {
	if err := c.Fix(); err != nil {
		return err
	}
	if err := c.ToWriter(w); err != nil {
		return err
	}
	return nil
}

func (c *Character) GetBytes() ([]byte, error) {
	var buf bytes.Buffer
	if err := c.ToWriter(&buf); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func (c *Character) GetCorrectBytes() ([]byte, error) {
	var buf bytes.Buffer
	bw := bufio.NewWriter(&buf)
	if err := c.ToWriterCorrect(bw); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
