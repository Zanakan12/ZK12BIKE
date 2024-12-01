package bikes

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
	"zk12ebike/internal/database"
	"zk12ebike/internal/home"
)

// Compteur global pour les noms de fichiers
const dirPath = "static/images/bike"

var mu sync.Mutex

func BikeListHandler(w http.ResponseWriter, r *http.Request) {
	// Parse les fichiers de template
	tmpl, err := template.ParseFiles("templates/base.html", "templates/navbar.html", "templates/bike_list.html")
	if err != nil {
		fmt.Println("Erreur lors du parsing des templates:", err)
		http.Error(w, "Erreur interne du serveur", http.StatusInternalServerError)
		return
	}
	allbike,_ := database.GetAllBikes()
	// Création des données à envoyer au template
	data := home.Pageinfo{
		Title: "Bike list",
		Page:  "bike-list",
		Bike:  allbike,
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
	counter, err := findMissingNumber(dirPath)
	if err != nil {
		fmt.Println("Erreur lors du chargement des noms de fichiers")
	}
	fileName := fmt.Sprintf("%d.jpg", counter)

	// Crée un dossier pour stocker les fichiers uploadés, s'il n'existe pas
	if err := os.MkdirAll(dirPath, os.ModePerm); err != nil {
		http.Error(w, "Erreur lors de la création du dossier de stockage", http.StatusInternalServerError)
		return
	}

	// Crée le fichier final avec le nom basé sur le compteur
	dst, err := os.Create(filepath.Join(dirPath, fileName))
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
	if r.Method == http.MethodPost {
		bike_type := r.FormValue("bike_type")
		size, _ := strconv.Atoi(r.FormValue("size"))
		motor_type := r.FormValue("motor_type")
		speed, _ := strconv.Atoi(r.FormValue("speed"))
		autonomy, _ := strconv.Atoi(r.FormValue("autonomy"))
		price, _ := strconv.Atoi(r.FormValue("price"))
		status := r.FormValue("status")
		battery, _ := strconv.Atoi(r.FormValue("battery"))
		// Ecrire dans la db
		filePath := "/"+dirPath+"/"+fileName
		database.SaveBikeToDB(filePath, bike_type, motor_type, status, float64(size), speed, autonomy, battery, float64(price))

	}
	http.Redirect(w, r, "/admin", http.StatusSeeOther)
}

func findMissingNumber(dirPath string) (int, error) {
	// Map pour stocker les entiers trouvés
	numbers := make(map[int]bool)

	// Parcourir les fichiers du répertoire
	err := filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Vérifier si ce n'est pas un répertoire
		if !info.IsDir() {
			// Extraire le nom du fichier sans extension
			fileName := strings.TrimSuffix(info.Name(), filepath.Ext(info.Name()))
			// Convertir en entier si possible
			if number, err := strconv.Atoi(fileName); err == nil {
				numbers[number] = true
			}
		}
		return nil
	})

	if err != nil {
		return 0, err
	}

	// Trouver le premier entier manquant
	for i := 0; ; i++ {
		if !numbers[i] {
			return i, nil
		}
	}
}
