package d2s

import (
	"github.com/vitalick/d2s/character"
	"log"
	"os"
)

func Parse(filepath string) (*character.Character, error) {
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
