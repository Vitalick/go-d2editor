package d2s

import (
	"fmt"
	"log"
	"os"
)

//Open returns a new Character for editing and viewing
func Open(filepath string) (*Character, error) {
	file, err := os.Open(filepath)
	if err != nil {
		log.Println("Error while opening .d2s file")
		return nil, err
	}

	defer file.Close()

	c, err := NewCharacter(file)

	if err != nil {
		log.Println("Error while parsing .d2s file")
		return nil, err
	}

	return c, nil
}

//Save will create *.d2s file in folder with Character struct
func Save(c *Character, folder string) error {
	if len(c.Name) == 0 {
		return ErrorBlankName
	}
	if len(folder) > 0 && folder[len(folder)-1:] != "/" {
		folder = folder + "/"
	}
	file, err := os.Create(fmt.Sprintf("%s%s.d2s", folder, c.Name))
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
