package database

func SaveUserToDB(name, email, password,role string) error {
	db := SetupDatabase()
	defer db.Close()

	_, err := db.Exec("INSERT INTO users (username, email, password,role) VALUES (?, ?, ?, ?)", name, email, password,role)
	return err
}

func SaveBikeToDB(image_path, bike_type, motor_type, status string, size float64, speed, autonomy, battery int, price float64) error {
	db := SetupDatabase()
	defer db.Close()

	// Vérifiez le nom correct de la table
	_, err := db.Exec(`
		INSERT INTO bikes 
		(image_path, bike_type, motor_type,  size, speed, autonomy, price, battery, status) 
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)`,
		image_path, bike_type, motor_type, size, speed, autonomy, price, battery, status, 
	)

	return err
}

