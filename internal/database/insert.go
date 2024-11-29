package database

func SaveUserToDB(name, email, password string) error {
	db := SetupDatabase()
	defer db.Close()

	_, err := db.Exec("INSERT INTO users (username, email, password) VALUES (?, ?, ?)", name, email, password)
	return err
}
