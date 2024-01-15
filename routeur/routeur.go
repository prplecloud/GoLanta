package routeur

import (
	"fmt"
	"golanta/controller"
	"net/http"
)

func Initserv() {

	css := http.FileServer(http.Dir("./assets/"))
	http.Handle("/static/", http.StripPrefix("/static/", css))

	http.HandleFunc("/creation", controller.Creation)
	http.HandleFunc("/profil", controller.Profil)
	http.HandleFunc("/treatment", controller.FormSubmission)

	fmt.Println("serveur ouvert sur le port 8080")
	http.ListenAndServe(":8080", nil)
}
