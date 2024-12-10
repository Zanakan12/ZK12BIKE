package home

import (
	"log"
	"net/http"
	"text/template"
	"zk12ebike/internal/cookies"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/index.html", "templates/base.html", "templates/navbar.html")
	if err != nil {
		log.Println("Erreur lors du chargement du template:", err)
		http.Error(w, "Erreur interne du serveur 1", http.StatusInternalServerError)
		return
	}
	session := cookies.GetCookie(w, r)
	username := "Biker"
	if session.Username != "" {
		username = session.Username
	}

	// Les données à passer au template
	data := Pageinfo{
		Title:    "Page d'accueil",
		Page:     "home",
		Username: username,
		Session:  session,
	}
	// Exécution du template
	if err := tmpl.ExecuteTemplate(w, "base.html", data); err != nil {
		log.Println("Erreur lors de l'exécution du template:", err)
		http.Error(w, "Erreur interne du serveur 2", http.StatusInternalServerError)
	}
}

func ContactHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/index.html", "templates/base.html", "templates/navbar.html", "templates/contact.html")
	if err != nil {
		log.Println("Erreur lors du chargement du template:", err)
		http.Error(w, "Erreur interne du serveur 1", http.StatusInternalServerError)
		return
	}
	session := cookies.GetCookie(w, r)
	username := "Biker"
	if session.Username != "" {
		username = session.Username
	}

	// Les données à passer au template
	data := Pageinfo{
		Title:    "Contact",
		Page:     "contact",
		Username: username,
		Session:  session,
	}
	// Exécution du template
	if err := tmpl.ExecuteTemplate(w, "base.html", data); err != nil {
		log.Println("Erreur lors de l'exécution du template:", err)
		http.Error(w, "Erreur interne du serveur 2", http.StatusInternalServerError)
	}
}

func PriceHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/index.html", "templates/base.html", "templates/navbar.html", "templates/price.html")
	if err != nil {
		log.Println("Erreur lors du chargement du template:", err)
		http.Error(w, "Erreur interne du serveur 1", http.StatusInternalServerError)
		return
	}
	session := cookies.GetCookie(w, r)
	username := "Biker"
	if session.Username != "" {
		username = session.Username
	}

	// Les données à passer au template
	data := Pageinfo{
		Title:    "Contact",
		Page:     "contact",
		Username: username,
		Session:  session,
	}
	// Exécution du template
	if err := tmpl.ExecuteTemplate(w, "base.html", data); err != nil {
		log.Println("Erreur lors de l'exécution du template:", err)
		http.Error(w, "Erreur interne du serveur 2", http.StatusInternalServerError)
	}
}

func LocationHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/index.html", "templates/base.html", "templates/navbar.html", "templates/locations.html")
	if err != nil {
		log.Println("Erreur lors du chargement du template:", err)
		http.Error(w, "Erreur interne du serveur 1", http.StatusInternalServerError)
		return
	}
	session := cookies.GetCookie(w, r)
	username := "Biker"
	if session.Username != "" {
		username = session.Username
	}

	// Les données à passer au template
	data := Pageinfo{
		Title:    "Location",
		Page:     "location",
		Username: username,
		Session:  session,
	}
	// Exécution du template
	if err := tmpl.ExecuteTemplate(w, "base.html", data); err != nil {
		log.Println("Erreur lors de l'exécution du template:", err)
		http.Error(w, "Erreur interne du serveur 2", http.StatusInternalServerError)
	}
}