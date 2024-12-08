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

func init() {
	database.SetupDatabase()
}

func Handler(w http.ResponseWriter, r *http.Request) {
	// Serve static files
	if strings.HasPrefix(r.URL.Path, "/static/") {
		http.StripPrefix("/static/", http.FileServer(http.Dir("static"))).ServeHTTP(w, r)
		return
	}

	// Route handling
	switch {
	case r.URL.Path == "/":
		home.HomeHandler(w, r)
	case r.URL.Path == "/register":
		users.RegisterHandler(w, r)
	case r.URL.Path == "/login":
		users.LoginHandler(w, r)
	case r.URL.Path == "/logout":
		users.LogoutHandler(w, r)
	case r.URL.Path == "/profile":
		users.ProfileHandler(w, r)
	case r.URL.Path == "/bike-list":
		bikes.BikeListHandler(w, r)
	case r.URL.Path == "/admin":
		users.AdminPanelHandler(w, r)
	case r.URL.Path == "/addtoshop":
		bikes.AddToCartHandler(w, r)
	case strings.HasPrefix(r.URL.Path, "/bike-detail/"):
		id, _ := strconv.Atoi(strings.TrimPrefix(r.URL.Path, "/bike-detail/"))
		bikes.BikeDetailHandler(w, r, id)
	case r.URL.Path == "/delete":
		bikes.DeleteBikeHandler(w, r)
	case r.URL.Path == "/upload":
		bikes.UploadFile(w, r)
	default:
		http.NotFound(w, r)
	}
}