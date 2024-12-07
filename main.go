package main

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"zk12ebike/internal/bikes"
	"zk12ebike/internal/database"
	"zk12ebike/internal/home"
	"zk12ebike/internal/users"

	"github.com/gin-gonic/gin"
)

const port = ":8080"

var (
	app *gin.Engine
)
func ini(){
	app = gin.New()
}

func main() {

	database.SetupDatabase()
	//Serve the static directory
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	//start the the server :
	http.HandleFunc("/", home.HomeHandler)
	http.HandleFunc("/register", users.RegisterHandler)
	http.HandleFunc("/login", users.LoginHandler)
	http.HandleFunc("/logout", users.LogoutHandler)
	http.HandleFunc("/profile", users.ProfileHandler)
	http.HandleFunc("/bike-list", bikes.BikeListHandler)
	http.HandleFunc("/admin", users.AdminPanelHandler)
	http.HandleFunc("/addtoshop", bikes.AddToCartHandler)
	http.HandleFunc("/bike-detail/", func(w http.ResponseWriter, r *http.Request) {
		// Récupère l'ID sur l'Url
		id, _ := strconv.Atoi(strings.TrimPrefix(r.URL.Path, "/bike-detail/"))
		bikes.BikeDetailHandler(w, r, id)
		return

	})
	// Action
	http.HandleFunc("/delete", bikes.DeleteBikeHandler)
	http.HandleFunc("/upload", bikes.UploadFile)
	// route to handlers make the route before the running server
	fmt.Println("The server is start on http://localhost" + port)
	if err := http.ListenAndServe(port, nil); err != nil {
		fmt.Println("Erreur lors du démarrage du serveur : ", err)
	}

}
