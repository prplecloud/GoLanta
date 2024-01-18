package backend

import (
	"encoding/json"
	"log"
	"math/rand"
	"os"
	"time"
)

var jsonfile = "data.json"

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

func GenerateRandomNumber() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(90000) + 10000
}

func ModifyChara(updatedChara Personnage) error {
	character, err := LoadChara()
	if err != nil {
		log.Fatal("log: retrieveArticles() error!\n", err)
	}
	for i, chara := range character {
		if chara.ID == updatedChara.ID {
			character[i] = updatedChara
		}
	}
	ChangeChara(character)
	return nil
}

func ChangeChara(character []Personnage) {

	data, errJSON := json.Marshal(character)
	if errJSON != nil {
		log.Fatal("log: addArticle()\t JSON Marshall error!\n", errJSON)
	}

	errWrite := os.WriteFile(jsonfile, data, 0666)
	if errWrite != nil {
		log.Fatal("log: addArticle()\t WriteFile error!\n", errWrite)
	}
}
