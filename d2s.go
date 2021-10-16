package d2editor

import (
	"encoding/json"
	"fmt"
	"os"
)

func beforeSave(c *Character, folder *string) error {
	if len(c.Name) == 0 {
		return errorBlankName
	}
	if len(*folder) > 0 && (*folder)[len(*folder)-1:] != "/" {
		*folder = *folder + "/"
	}
	return nil
}

//Open returns a new Character from .d2s file
func Open(filepath string) (*Character, error) {
	file, err := os.Open(filepath)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error while opening .d2s file")
		return nil, err
	}

	defer file.Close()

	c, err := NewCharacter(file)

	if err != nil {
		fmt.Fprintln(os.Stderr, "Error while parsing .d2s file")
		return nil, err
	}

	return c, nil
}

//Save will create .d2s file in folder with Character struct
func Save(c *Character, folder string) error {
	if err := beforeSave(c, &folder); err != nil {
		fmt.Fprintln(os.Stderr, errorBlankName)
		return err
	}
	file, err := os.Create(fmt.Sprintf("%s%s.d2s", folder, c.Name))
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error while creating .d2s file")
		return err
	}

	defer file.Close()

	err = c.ToWriterCorrect(file)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error while writing buffer file")
		return err
	}
	return nil
}

//OpenJSON returns a new Character from json file
func OpenJSON(filepath string) (*Character, error) {
	file, err := os.Open(filepath)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error while opening json file")
		return nil, err
	}

	defer file.Close()

	c, err := NewEmptyCharacter(97)

	if err != nil {
		fmt.Fprintln(os.Stderr, "Error while create empty Character")
		return nil, err
	}

	d := json.NewDecoder(file)
	err = d.Decode(&c)

	if err != nil {
		fmt.Fprintln(os.Stderr, "Error while parsing json file")
		return nil, err
	}

	return c, nil
}

//SaveJSON will create *.d2s file in folder with Character struct
func SaveJSON(c *Character, folder string) error {
	if err := beforeSave(c, &folder); err != nil {
		fmt.Fprintln(os.Stderr, errorBlankName)
		return err
	}
	file, err := os.Create(fmt.Sprintf("%s%s.json", folder, c.Name))
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error while creating json file")
		return err
	}

	defer file.Close()

	d := json.NewEncoder(file)
	err = d.Encode(&c)

	if err != nil {
		fmt.Fprintln(os.Stderr, "Error while writing json file")
		return err
	}
	return nil
}
