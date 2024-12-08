package handler

import (
	"net/http"
	"strconv"
	"strings"
	"zk12ebike/internal/bikes"
	"zk12ebike/internal/database"
	"zk12ebike/internal/home"
	"zk12ebike/internal/users"
)


// Cette fonction exportée est ce que Vercel attend comme point d'entrée
func Handler(w http.ResponseWriter, r *http.Request) {
	// Serveur de base pour ton application web
	database.SetupDatabase() // Assure-toi que la base de données est prête

	// Serve les fichiers statiques
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Définir les gestionnaires pour chaque route
	http.HandleFunc("/", home.HomeHandler)
	http.HandleFunc("/register", users.RegisterHandler)
	http.HandleFunc("/login", users.LoginHandler)
	http.HandleFunc("/logout", users.LogoutHandler)
	http.HandleFunc("/profile", users.ProfileHandler)
	http.HandleFunc("/bike-list", bikes.BikeListHandler)
	http.HandleFunc("/admin", users.AdminPanelHandler)
	http.HandleFunc("/addtoshop", bikes.AddToCartHandler)

	// Détail d'un vélo par son ID
	http.HandleFunc("/bike-detail/", func(w http.ResponseWriter, r *http.Request) {
		// Récupère l'ID sur l'URL
		id, _ := strconv.Atoi(strings.TrimPrefix(r.URL.Path, "/bike-detail/"))
		bikes.BikeDetailHandler(w, r, id)
	})

	// Routes pour les actions
	http.HandleFunc("/delete", bikes.DeleteBikeHandler)
	http.HandleFunc("/upload", bikes.UploadFile)

}
