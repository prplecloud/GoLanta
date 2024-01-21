package controller

import (
	"encoding/json"
	"fmt"
	"golanta/backend"
	"golanta/templates"
	"io/fs"
	"net/http"
	"os"
	"strconv"
)

func Profil(w http.ResponseWriter, r *http.Request) {
	backend.Chara, _ = backend.LoadChara()
	fmt.Println(backend.Chara)
	fmt.Println(len(backend.Chara))
	templates.Temp.ExecuteTemplate(w, "profil", backend.Chara)
}

func Creation(w http.ResponseWriter, r *http.Request) {
	templates.Temp.ExecuteTemplate(w, "creation", nil)
}

func Modify(w http.ResponseWriter, r *http.Request) {
	templates.Temp.ExecuteTemplate(w, "modify", nil)
}

func FormSubmission(w http.ResponseWriter, r *http.Request) {

	nomFichier := "data.json"

	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Erreur lors de l'analyse du formulaire", http.StatusInternalServerError)
		return
	}

	id := backend.GenerateRandomNumber()

	form := backend.Personnage{
		ID:             id,
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

func Delete(w http.ResponseWriter, r *http.Request) {

	var err error

	if r.Method == http.MethodPost {
		backend.Chara, err = backend.LoadChara()
		if err != nil {
			http.Error(w, fmt.Sprintf("Erreur du JSON : %s", err), http.StatusInternalServerError)
			return
		}
		err = r.ParseForm()
		if err != nil {
			http.Error(w, fmt.Sprintf("Erreur de parse du formulaire : %s", err), http.StatusInternalServerError)
			return
		}

		deleteIDStr := r.FormValue("article_id")
		if deleteIDStr == "" {
			http.Error(w, "ID ne peut pas être une chaîne vide", http.StatusBadRequest)
			return
		}

		deleteID, err := strconv.Atoi(deleteIDStr)
		if err != nil {
			http.Error(w, fmt.Sprintf("Erreur de conversion string to int : %s", err), http.StatusInternalServerError)
			return
		}

		backend.RemoveCharater(deleteID)
	}

	http.Redirect(w, r, "http://localhost:8080/profil", http.StatusSeeOther)
}

func ModifyCharaHandler(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()

	CharaId := r.FormValue("character")
	CharaIdInt, _ := strconv.Atoi(CharaId)
	updatedChara := backend.GetCharacter(CharaIdInt)

	fmt.Println(updatedChara.Nom)

	w.WriteHeader(http.StatusOK)

	fmt.Println("crampté")

	err := templates.Temp.ExecuteTemplate(w, "modify", backend.Chara)
	if err != nil {
		http.Error(w, "Erreur d'affichage du template", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "http://localhost:8080/profil", http.StatusSeeOther)
}
