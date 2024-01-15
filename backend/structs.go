package backend

type Personnage struct {
	Id             int    `json: "id"`
	Nom            string `json: "nom"`
	Sexe           string `json: "sexe"`
	CouleurCheveux string `json: "couleurcheveux"`
	Equipe         string `json: "equipe"`
}
