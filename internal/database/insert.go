package database

func SaveUserToDB(name, email, password, role, profile_image string) error {
	db := SetupDatabase()
	defer db.Close()

	_, err := db.Exec("INSERT INTO users (username, email, password,role, profile_image) VALUES (?, ?, ?, ?, ?)", name, email, password, role, profile_image)
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

func SaveShopToDB(user_id, bike_id int, bike_type, image_path, status string, price, size float64, total int) error {
	db := SetupDatabase()
	defer db.Close()
	
	_, err := db.Exec(`
		INSERT INTO shop 
		(user_id, image_path, bike_id, bike_type, status, price, size, total) 
		VALUES (?, ?, ?, ?, ?, ?, ?, ?)`,
		user_id, image_path, bike_id, bike_type, status, price, size, total,
	)

	return err
}
