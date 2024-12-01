package bikes

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"sync"
	"zk12ebike/internal/home"
)

// Compteur global pour les noms de fichiers
var counter int
var mu sync.Mutex

func BikeListHandler(w http.ResponseWriter, r *http.Request) {
	// Parse les fichiers de template
	tmpl, err := template.ParseFiles("../templates/base.html", "../templates/navbar.html", "../templates/bike_list.html")
	if err != nil {
		fmt.Println("Erreur lors du parsing des templates:", err)
		http.Error(w, "Erreur interne du serveur", http.StatusInternalServerError)
		return
	}

	// Création des données à envoyer au template
	data := home.Pageinfo{
		Title: "Bike list",
		Page:  "bike-list",
	}
	fmt.Println(data)
	// Exécution du template
	if err := tmpl.ExecuteTemplate(w, "base.html", data); err != nil {
		log.Println("Erreur lors de l'exécution du template:", err)
		http.Error(w, "Erreur interne du serveur 2", http.StatusInternalServerError)
	}
}


// Fonction pour gérer l'upload du fichier
func UploadFile(w http.ResponseWriter, r *http.Request) {
	// Limite de taille du fichier (ici 10 MB)
	const MAX_UPLOAD_SIZE = 10 * 1024 * 1024 // 10MB
	r.Body = http.MaxBytesReader(w, r.Body, MAX_UPLOAD_SIZE)

	// Vérifie si le formulaire a été soumis
	if err := r.ParseMultipartForm(MAX_UPLOAD_SIZE); err != nil {
		http.Error(w, "Le fichier est trop volumineux ou problème avec l'upload", http.StatusBadRequest)
		return
	}

	// Récupère le fichier depuis le formulaire
	file, _, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Erreur lors du téléchargement du fichier", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	// Verrouille l'accès au compteur pour garantir que le nom de fichier est unique
	mu.Lock()
	defer mu.Unlock()

	// Utilise le compteur pour créer un nom de fichier basé sur un chiffre
	fileName := fmt.Sprintf("%d.jpg", counter)
	counter++

	// Crée un dossier pour stocker les fichiers uploadés, s'il n'existe pas
	if err := os.MkdirAll("uploads", os.ModePerm); err != nil {
		http.Error(w, "Erreur lors de la création du dossier de stockage", http.StatusInternalServerError)
		return
	}

	// Crée le fichier final avec le nom basé sur le compteur
	dst, err := os.Create(filepath.Join("uploads", fileName))
	if err != nil {
		http.Error(w, "Erreur lors de la sauvegarde du fichier", http.StatusInternalServerError)
		return
	}
	defer dst.Close()

	// Copie le contenu du fichier téléchargé dans le fichier créé
	_, err = dst.ReadFrom(file)
	if err != nil {
		http.Error(w, "Erreur lors de la sauvegarde du fichier", http.StatusInternalServerError)
		return
	}

	// Répond à l'utilisateur
	fmt.Fprintf(w, "Fichier téléchargé avec succès sous le nom : %s", fileName)
}
