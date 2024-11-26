package database

import "database/sql"

func createUsersTable(db *sql.DB) {
	createTableSQL := `CREATE TABLE IF NOT EXISTS users (
		user_id INTEGER NOT NULL UNIQUE,
		name TEXTE NOT NULL,
		email TEXTE NOT NULL,
		password TEXTE NOT NULL,
		created_at TIMESAMP DEFAULT CURRENT_TIMESAMP
	);`

	executeSQL(db, createTableSQL)
}

func createOrderTable(db *sql.DB) {
	createTableSQL := `CREATE TABLE IF NOT EXISTS order (
		id INTEGER NOT NULL UNIQUE,
		users_id INTEGER NOT NULL UNIQUE,
		bike_id	INTEGER NOT NULL UNIQUE,
		bike_type TEXT NOT NULL,
		start_date TIMESAMP DEFAULT CURRENT_TIMESAMP
		end_date	TIMESAMP DEFAULT CURRENT_TIMESAMP
		total_price float64
		status	TEXT NOT NULL,
		created_at TIMESAMP DEFAULT CURRENT_TIMESAMP
	);`

	executeSQL(db, createTableSQL)
}

func createBikesTable(db *sql.DB) {
	createTableSQL := `CREATE TABLE IF NOT EXISTS bikes (
		id INTEGER NOT NULL UNIQUE,
		name TEXTE NOT NULL,
		bike_type TEXT NOT NULL,
		price FLOAT64,
		status TEXT NOT NULL,
		created_at TIMESAMP DEFAULT CURRENT_TIMESAMP
	);`

	executeSQL(db, createTableSQL)
}