package home

import "zk12ebike/internal/cookies"


type Pageinfo struct {
	Title string
	Page string
	Logo string
	Username string
	Session cookies.Session
}
