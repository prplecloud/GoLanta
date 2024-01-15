package backend

import (
	"encoding/json"
	"os"
)

func LoadChara() ([]Personnage, error) {

	fileData, err := os.ReadFile("data.json")
	if err != nil {
		return nil, err
	}

	var forms []Personnage

	if len(fileData) != 0 {
		err = json.Unmarshal(fileData, &forms)
		if err != nil {
			return nil, err
		}
	}

	return forms, nil
}
