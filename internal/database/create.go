package database

import "database/sql"

// Créer la table des utilisateurs
func createUsersTable(db *sql.DB) {
	createTableSQL := `CREATE TABLE IF NOT EXISTS users (
		user_id INTEGER PRIMARY KEY AUTOINCREMENT,
		username TEXT NOT NULL,
		email TEXT NOT NULL UNIQUE,
		password TEXT NOT NULL,
		role TEXT NOT NULL,
		profile_image TEXT,
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
		image_path TEXT NOT NULL,
		bike_type TEXT NOT NULL,
		motor_type TEXT NOT NULL,
		size REAL,
		speed INTEGER NOT NULL,
		autonomy INTERGER NOT NULL,
		price REAL NOT NULL,
		battery INTEGER NOT NULL,
		status TEXT NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);`

	executeSQL(db, createTableSQL)
}

func createShopTable(db *sql.DB) {
	createTableSQL := `CREATE TABLE IF NOT EXISTS shop (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		image_path TEXT,
		user_id INTEGER NOT NULL,
		bike_id INTEGER NOT NULL,
		bike_type TEXT NOT NULL,
		status TEXT,
		price REAL NOT NULL,
		size REAL NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		total INTEGER NOT NULL,
		FOREIGN KEY (user_id) REFERENCES users(user_id),
		FOREIGN KEY (bike_id) REFERENCES bikes(id)

	);`

	executeSQL(db, createTableSQL)
}
