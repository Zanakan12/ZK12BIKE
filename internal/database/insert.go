package database

func SaveUserToDB(name, email, password,role string) error {
	db := SetupDatabase()
	defer db.Close()

	_, err := db.Exec("INSERT INTO users (username, email, password,role) VALUES (?, ?, ?, ?)", name, email, password,role)
	return err
}

