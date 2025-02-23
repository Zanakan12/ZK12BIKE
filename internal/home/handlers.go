package home

import (
	"log"
	"net/http"
	"text/template"
	"zk12ebike/internal/cookies"
	"encoding/json"
	"strings"
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



// We’ll use this as our “database”
var items = []string{
    "Apple", "Banana", "Orange", "Papaya", 
    "Mango", "Melon", "Peach", "Blueberry", 
    "Strawberry", "Watermelon",
}


// handleSearch processes the ?query= parameter and returns filtered items
func HandleSearch(w http.ResponseWriter, r *http.Request) {
    query := r.URL.Query().Get("query")
    if query == "" {
        // No query => return empty result or everything, up to you
        json.NewEncoder(w).Encode([]string{})
        return
    }

    // Filter items
    var results []string
    for _, item := range items {
        if strings.Contains(strings.ToLower(item), strings.ToLower(query)) {
            results = append(results, item)
        }
    }

    // Return filtered items as JSON
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(results)
}
