package backend

import (
	"encoding/json"
	"math/rand"
	"os"
	"time"
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

func GenerateID() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(9000) + 1000
}

func CharaID(filePath string) ([]int, error) {

	charaID := []int{}

	fileContent, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	var character []Personnage

	err = json.Unmarshal(fileContent, &character)
	if err != nil {
		return nil, err
	}

	for _, chara := range character {
		charaID = append(charaID, chara.ID)
	}
	return charaID, nil
}
