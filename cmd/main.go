package main

import (
	"fmt"
	"net/http"
	"zk12ebike/internal/home"
)

const port = ":8080"

func main() {

	//Serve the static directory
	fs := http.FileServer(http.Dir("../static"))
    http.Handle("/static/", http.StripPrefix("/static/", fs))

	//start the the server :
	http.HandleFunc("/", home.HomeHandler)
	//http.HandleFunc("/register",users.RegisterHandler)

	// route to handlers make the route before the running server
	fmt.Println("The server is start on http://localhost" + port)
	if err := http.ListenAndServe(port, nil); err != nil {
		fmt.Println("Erreur lors du démarrage du serveur : ", err)
	}

	fmt.Println("test 1")

}
