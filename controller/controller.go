package controller

import (
	"encoding/json"
	"fmt"
	"golanta/backend"
	"golanta/templates"
	"io/fs"
	"net/http"
	"os"
)

func Home(w http.ResponseWriter, r *http.Request) {
	templates.Temp.ExecuteTemplate(w, "home", nil)
}

func Profil(w http.ResponseWriter, r *http.Request) {
	templates.Temp.ExecuteTemplate(w, "profil", nil)
}

func Creation(w http.ResponseWriter, r *http.Request) {
	templates.Temp.ExecuteTemplate(w, "creation", nil)
}

func FormSubmission(w http.ResponseWriter, r *http.Request) {

	nomFichier := "data.json"

	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Erreur lors de l'analyse du formulaire", http.StatusInternalServerError)
		return
	}

	// Créer une nouvelle instance de Form à partir des données du formulaire
	form := backend.Personnage{
		Nom:            r.FormValue("nom"),
		Sexe:           r.FormValue("sexe"),
		CouleurCheveux: r.FormValue("couleurcheveux"),
		Equipe:         r.FormValue("equipe"),
	}

	dataForms, errForms := backend.LoadChara()
	if errForms != nil {
		http.Error(w, fmt.Sprintf("Erreur lors de l'ouverture du fichier : %v", errForms), http.StatusInternalServerError)
		return
	}

	// Ajouter la nouvelle forme à la liste
	dataForms = append(dataForms, form)

	dataWrite, errWrite := json.Marshal(dataForms)
	if errWrite != nil {
		http.Error(w, fmt.Sprintf("Erreur lors du marshal du fichier : %v", errWrite), http.StatusInternalServerError)
		return
	}

	errWriteFile := os.WriteFile(nomFichier, dataWrite, fs.FileMode(0644))
	if errWriteFile != nil {
		http.Error(w, fmt.Sprintf("Erreur lors de l'écriture du fichier : %v", errWriteFile), http.StatusInternalServerError)
		return
	}

	fmt.Println("Ajouté avec succès dans le JSon")
	http.Redirect(w, r, "http://localhost:8080/profil", http.StatusSeeOther)
}
