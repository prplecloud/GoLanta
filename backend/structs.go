package backend

type Personnage struct {
	ID             int    `json: "ID"`
	Nom            string `json: "nom"`
	Sexe           string `json: "sexe"`
	CouleurCheveux string `json: "couleurcheveux"`
	Equipe         string `json: "equipe"`
}

type CharaData struct {
	Charakter []Personnage `json: "charakter"`
}

var Chara []Personnage
