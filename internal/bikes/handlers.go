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
	"zk12ebike/internal/cookies"
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

	session := cookies.GetCookie(w, r)
	allbike, _ := database.GetAllBikes()

	username := "Biker"
	if session.Username != "" {
		username = session.Username
	}
	// Création des données à envoyer au template
	data := home.Pageinfo{
		Title:    "Bike list",
		Page:     "bike-list",
		Username: username,
		Session:  session,
		Bike:     allbike,
	}

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
		filePath := "/" + dirPath + "/" + fileName
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

func DeleteBikeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		bike_id, _ := strconv.Atoi(r.FormValue("bike_id"))
		action := r.FormValue("delete")
		fildPath := r.FormValue("fildPath")

		session := cookies.GetCookie(w, r)
		if session.UserID == 0 {
			http.Redirect(w, r, "/", http.StatusSeeOther)
		}

		if action == "delete" {
			database.DeleteBike(bike_id, fildPath)
		}
		http.Redirect(w, r, "/admin", http.StatusSeeOther)
	}
}

func BikeDetailHandler(w http.ResponseWriter, r *http.Request, id int) {
	tmpl, err := template.ParseFiles("templates/base.html", "templates/navbar.html", "templates/bike_detail.html")
	if err != nil {
		log.Println("Error during parse template")
	}
	session := cookies.GetCookie(w, r)
	oneBike, _ := database.GetOneBike(id)
	// Création des données à envoyer au template
	data := home.Pageinfo{
		Title:   "Bike Detail",
		Page:    "bike-detail",
		Bike:    oneBike,
		Session: session,
	}
	if err := tmpl.ExecuteTemplate(w, "base.html", data); err != nil {
		log.Println("Erreur lors de l'exécution du template:", err)
		http.Error(w, "Erreur interne du serveur : BikeDetailHandler", http.StatusInternalServerError)
	}
}

func AddToCartHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		user_id, _ := strconv.Atoi(r.FormValue("user_id"))
		bike_id, _ := strconv.Atoi(r.FormValue("bike_id"))
		bike_type := r.FormValue("bike_type")
		price, _ := strconv.Atoi(r.FormValue("price"))
		size, _ := strconv.Atoi(r.FormValue("size"))
		bikes, _ := database.GetOneBike(bike_id)
		image_path := ""
		status := ""
		if len(bikes) > 0 {
			bike := bikes[0]
			image_path = bike.ImagePath
			status = bike.Status
		}

		if user_id == 0 {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
		}
		total, verif := database.VerifBikeId(user_id, bike_id)
		if verif {
			total++
			database.UpdateShop(user_id, bike_id, total)
		} else {
			err := database.SaveShopToDB(user_id, bike_id, bike_type, image_path, status, float64(price), float64(size), total)
			if err != nil {
				log.Println("Error During save form in the db", err)
			}
		}

		http.Redirect(w, r, "/bike-detail/"+r.FormValue("bike_id"), http.StatusSeeOther)
	}

}

func UpdateStatusHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		bike_id, _ := strconv.Atoi(r.FormValue("bike_id"))
		status := r.FormValue("status")
		err := database.UpdateStatus(bike_id, status)
		if err != nil {
			log.Println("erreur lors du  changement du status: %v", err)
		}
	}
	http.Redirect(w, r, "/admin", http.StatusSeeOther)
}
