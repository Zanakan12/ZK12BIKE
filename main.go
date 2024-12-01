package main

import (
	"fmt"
	"net/http"
	"zk12ebike/internal/bikes"
	"zk12ebike/internal/database"
	"zk12ebike/internal/home"
	"zk12ebike/internal/users"
)

const port = ":8080"

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
	http.HandleFunc("/bike-list",bikes.BikeListHandler)
	http.HandleFunc("/admin",users.AdminPanelHandler)
	http.HandleFunc("/upload", bikes.UploadFile)
	// route to handlers make the route before the running server
	fmt.Println("The server is start on http://localhost" + port)
	if err := http.ListenAndServe(port, nil); err != nil {
		fmt.Println("Erreur lors du démarrage du serveur : ", err)
	}

}
