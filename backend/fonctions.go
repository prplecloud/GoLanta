package backend

import (
	"encoding/json"
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
		return err
	}

	for i, chara := range character {
		if chara.ID == updatedChara.ID {
			character[i] = updatedChara
		}
	}

	if err := ChangeChara(character); err != nil {
		return err
	}

	return nil
}

func ChangeChara(character []Personnage) error {

	data, err := json.MarshalIndent(character, "   ", " ")
	if err != nil {
		return err
	}

	err = os.WriteFile(jsonfile, data, 0666)
	if err != nil {
		return err
	}

	return nil
}

func RemoveCharater(id int) {
	var new []Personnage
	for _, i := range Chara {
		if i.ID != id {
			new = append(new, i)
		}
	}
	Chara = new
	ChangeChara(Chara)
}

func GetCharacter(id int) Personnage {
	Chara, _ := LoadChara()
	for _, i := range Chara {
		if i.ID == id {
			return i
		}
	}
	return Chara[0]
}
