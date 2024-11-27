package home

import (
	"log"
	"net/http"
	"text/template"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("../templates/index.html", "../templates/base.html", "../templates/navbar.html")
	if err != nil {
		log.Println("Erreur lors du chargement du template:", err)
		http.Error(w, "Erreur interne du serveur 1", http.StatusInternalServerError)
		return
	}

	// Les données à passer au template
	data := struct {
		Title string
	}{
		Title: "Bienvenue sur Mon Site",
	}

	// Exécution du template
	if err := tmpl.ExecuteTemplate(w, "base.html",data); err != nil {
		log.Println("Erreur lors de l'exécution du template:", err)
		http.Error(w, "Erreur interne du serveur 2", http.StatusInternalServerError)
	}
}
