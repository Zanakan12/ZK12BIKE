package home

import (
	"zk12ebike/internal/cookies"
	"zk12ebike/internal/database"
)

type Pageinfo struct {
	Title    string
	Page     string
	Logo     string
	Username string
	Cart	 int
	AllCart  float64
	Session  cookies.Session
	Bike     []database.Bike
	BikeShop []database.BikeShop
}
