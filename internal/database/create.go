package database

import "database/sql"

// Créer la table des utilisateurs
func createUsersTable(db *sql.DB) {
	createTableSQL := `CREATE TABLE IF NOT EXISTS users (
		user_id INTEGER PRIMARY KEY AUTOINCREMENT,
		username TEXT NOT NULL,
		email TEXT NOT NULL UNIQUE,
		password TEXT NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);`

	executeSQL(db, createTableSQL)
}

// Créer la table des commandes
func createOrdersTable(db *sql.DB) {
	createTableSQL := `CREATE TABLE IF NOT EXISTS orders (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		user_id INTEGER NOT NULL,
		bike_id INTEGER NOT NULL,
		bike_type TEXT NOT NULL,
		start_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		end_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		total_price REAL,
		status TEXT NOT NULL DEFAULT 'Disponible',
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		FOREIGN KEY (user_id) REFERENCES users(user_id),
		FOREIGN KEY (bike_id) REFERENCES bikes(id)
	);`

	executeSQL(db, createTableSQL)
}

// Créer la table des vélos
func createBikesTable(db *sql.DB) {
	createTableSQL := `CREATE TABLE IF NOT EXISTS bikes (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		username TEXT NOT NULL,
		bike_type TEXT NOT NULL,
		price REAL,
		status TEXT NOT NULL DEFAULT 'DISPONIBLE',
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);`

	executeSQL(db, createTableSQL)
}
