package d2s

import (
	"github.com/vitalick/d2s/character"
	"log"
	"os"
)

//Open returns a new Character for editing and viewing
func Open(filepath string) (*character.Character, error) {
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatalln("Error while opening .d2s file")
		return nil, err
	}

	defer file.Close()

	c, err := character.NewCharacter(file)

	if err != nil {
		log.Fatalln("Error while parsing .d2s file")
		return nil, err
	}

	return c, nil
}

//Save will create *.d2s file by filepath with Character struct
func Save(c *character.Character, filepath string) error {
	file, err := os.Create(filepath)
	if err != nil {
		log.Fatalln("Error while creating .d2s file")
		return err
	}

	defer file.Close()

	err = c.ToWriterCorrect(file)
	if err != nil {
		log.Fatalln("Error while writing buffer file")
		return err
	}
	return nil
}
