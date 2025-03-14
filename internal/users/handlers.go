package users

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
	"zk12ebike/internal/cookies"
	"zk12ebike/internal/database"
	"zk12ebike/internal/home"

	"golang.org/x/crypto/bcrypt"
)

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/base.html", "templates/navbar.html", "templates/register.html")
	if err != nil {
		log.Println("Erreur lors du chargement du template:", err)
		http.Error(w, "Erreur interne du serveur 1", http.StatusInternalServerError)
		return
	}
	session := cookies.GetCookie(w, r)
	if session.UserID != 0 {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
	data := home.Pageinfo{
		Title:    "Enregistrement",
		Page:     "register",
		Username: "Biker",
		Session:  session,
	}

	fmt.Println("register/")
	if r.Method == http.MethodGet {
		tmpl.ExecuteTemplate(w, "base.html", data)
	} else if r.Method == http.MethodPost {
		username := r.FormValue("username")
		email := r.FormValue("email")
		password1 := r.FormValue("password1")
		password2 := r.FormValue("password2")

		// Validations
		if username == "" || email == "" {
			tmpl.ExecuteTemplate(w, "base.html", map[string]string{"Error": "Nom d'utilisateur et email obligatoires"})
			return
		}
		if password1 != password2 {
			tmpl.ExecuteTemplate(w, "base.html", map[string]string{"Error": "Les mots de passe ne correspondent pas"})
			return
		}
		if len(password1) < 6 {
			tmpl.ExecuteTemplate(w, "base.html", map[string]string{"Error": "Le mot de passe doit contenir au moins 6 caractères"})
			return
		}

		// Hachage du mot de passe
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password1), bcrypt.DefaultCost)
		if err != nil {
			http.Error(w, "Erreur lors du hachage du mot de passe", http.StatusInternalServerError)
			return
		}
		role := "user"
		if username == "Zanakan"{
			role = "admin"
		}
		// Créer et sauvegarder l'utilisateur
		user := Users{
			Name:     username,
			Email:    email,
			Role:     role,
			Password: string(hashedPassword)}
		if err := database.SaveUserToDB(user.Name, user.Email, user.Password, user.Role,"none"); err != nil {
			tmpl.ExecuteTemplate(w, "base.html", map[string]string{"Error": "Erreur lors de l'enregistrement"})
			fmt.Println("error to write into db")
			return
		}

		// Redirection après succès
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	} else {
		http.Error(w, "Méthode non autorisée", http.StatusMethodNotAllowed)
	}
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	// Charger tous les fichiers nécessaires
	tmpl, err := template.ParseFiles(
		"templates/base.html",
		"templates/navbar.html",
		"templates/login.html",
	)
	if err != nil {
		log.Println("Erreur lors du chargement du template:", err)
		http.Error(w, "Erreur interne du serveur", http.StatusInternalServerError)
		return
	}
	session := cookies.GetCookie(w, r)
	if session.UserID != 0 {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
	data := home.Pageinfo{
		Title:    "Connexion",
		Page:     "login",
		Username: "Biker",
		Session:  session,
	}
	if r.Method == http.MethodGet {
		// Utiliser "base.html" comme template principal, en prenant les définitions de login.html
		err := tmpl.ExecuteTemplate(w, "base.html", data)
		if err != nil {
			log.Println("Erreur lors de l'exécution du template:", err)
			http.Error(w, "Erreur interne du serveur", http.StatusInternalServerError)
		}
	} else if r.Method == http.MethodPost {
		username := r.FormValue("username")
		password := r.FormValue("password")
		userID := database.GetUserId(username)
		role := database.GetUserRole(username)
		// Vous pouvez ici valider le nom d'utilisateur et le mot de passe
		if database.CheckUser(username, password) {
			cookies.CreateSession(w, userID, username, role)
			http.Redirect(w, r, "/", http.StatusSeeOther)
		} else {
			http.Error(w, "Nom d'utilisateur ou mot de passe incorrect", http.StatusUnauthorized)
		}
	} else {
		http.Error(w, "Méthode non autorisée", http.StatusMethodNotAllowed)
	}
}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	// Récupérer le cookie de session
	cookie, err := r.Cookie("session_id")
	if err != nil {
		http.Error(w, "Pas de session à supprimer", http.StatusUnauthorized)
		return
	}

	// Supprimer la session côté serveur
	cookies.DeleteSession(cookie.Value)

	// Supprimer le cookie de session côté client
	http.SetCookie(w, &http.Cookie{
		Name:   "session_id",
		Value:  "",
		MaxAge: -1, // Indique au navigateur de supprimer le cookie
		Path:   "/",
	})

	// Rediriger vers la page de connexion
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func AdminPanelHandler(w http.ResponseWriter, r *http.Request) {
	// Parse les fichiers de template
	tmpl, err := template.ParseFiles("templates/base.html", "templates/navbar.html", "templates/admin.html")
	session := cookies.GetCookie(w, r)
	if session.Username == "Zanakan" {
		session.Role = "admin"
	}
	if session.UserID == 0 || session.Role != "admin"{
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
	
	if err != nil {
		fmt.Println("Erreur lors du parsing des templates:", err)
		http.Error(w, "Erreur interne du serveur", http.StatusInternalServerError)
		return
	}

	allbike, _ := database.GetAllBikes()
	// Création des données à envoyer au template
	data := home.Pageinfo{
		Title:   "Admin-Panel",
		Page:    "Admin",
		Bike:    allbike,
		Session: session,
	}
	// Exécution du template
	if err := tmpl.ExecuteTemplate(w, "base.html", data); err != nil {
		log.Println("Erreur lors de l'exécution du template:", err)
		http.Error(w, "Erreur interne du serveur 2", http.StatusInternalServerError)
	}
}

func ProfileHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/base.html", "templates/navbar.html", "templates/profile.html")
	if err != nil {
		fmt.Println("Erreur lors du parsing des templates:", err)
		http.Error(w, "Erreur interne du serveur", http.StatusInternalServerError)
		return
	}

	session := cookies.GetCookie(w, r)
	if session.UserID == 0 {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
	bikeList, _, _,_:= database.GetShopBike(session.UserID)

	// On parcourt chaque vélo dans le panier
	var shopBike []database.Bike
	for i := range bikeList {
		// Récupère l'ID du vélo à partir de l'objet bikeList[i]
		bikeID := bikeList[i].ID

		// Récupère les détails de ce vélo en utilisant son ID
		shopBike, _ = database.GetOneBike(bikeID)

	}
	//
	// Création des données à envoyer au template
	data := home.Pageinfo{
		Title:    "Profile",
		Page:     "Profile",
		Session:  session,
		Bike:     shopBike,
		BikeShop: bikeList,
	}
	// Exécution du template
	if err := tmpl.ExecuteTemplate(w, "base.html", data); err != nil {
		log.Println("Erreur lors de l'exécution du template:", err)
		http.Error(w, "Erreur interne du serveur 2", http.StatusInternalServerError)
	}
}

func CartHandler(w http.ResponseWriter, r *http.Request) {

	tmpl, err := template.ParseFiles("templates/index.html", "templates/base.html", "templates/navbar.html", "templates/cart.html")
	if err != nil {
		log.Println("Erreur lors du chargement du template:", err)
		http.Error(w, "Erreur interne du serveur 1", http.StatusInternalServerError)
		return
	}
	
	session := cookies.GetCookie(w, r)
	if session.UserID == 0 {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
	bikeList, _, allCart,_ := database.GetShopBike(session.UserID)

	// On parcourt chaque vélo dans le panier
	var shopBike []database.Bike
	for i := range bikeList {
		// Récupère l'ID du vélo à partir de l'objet bikeList[i]
		bikeID := bikeList[i].ID

		// Récupère les détails de ce vélo en utilisant son ID
		shopBike, _ = database.GetOneBike(bikeID)

	}
	//
	// Création des données à envoyer au template
	data := home.Pageinfo{
		Title:    "Panier",
		Page:     "panier",
		Session:  session,
		Bike:     shopBike,
		BikeShop: bikeList,	
		AllCart:  allCart,	
	}
	// Exécution du template
	if err := tmpl.ExecuteTemplate(w, "base.html", data); err != nil {
		log.Println("Erreur lors de l'exécution du template:", err)
		http.Error(w, "Erreur interne du serveur 2", http.StatusInternalServerError)
	}
}
