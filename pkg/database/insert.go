package database

import "fmt"

func RequestInsert(userID int, username string) error {
	db := SetupDatabase()
	defer db.Close()

	tx, err := db.Begin()
	if err != nil {
		return fmt.Errorf("error starting transaction: %v", err)
	}

	insertSQL := `INSERT INTO requests (user_id, user_username) VALUES (?, ?)`
	_, err = tx.Exec(insertSQL, userID, username)

	if err != nil {
		tx.Rollback()
		return fmt.Errorf("error executing statement: %v", err)
	}

	if err = tx.Commit(); err != nil {
		return fmt.Errorf("error committing transaction: %v", err)
	}

	return nil
}
