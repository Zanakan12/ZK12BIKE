package database

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
	"golang.org/x/crypto/bcrypt"
)

func SetupDatabase() *sql.DB {
	db, err := sql.Open("sqlite3", "./db.sqlite")
	if err != nil {
		log.Fatal(err)
	}
	createUsersTable(db)
	createBikesTable(db)
	createOrdersTable(db)
	return db
}

func executeSQL(db *sql.DB, sql string) {
	// Prepare the SQL statement for execution
	statement, err := db.Prepare(sql)
	if err != nil {
		log.Fatal(err) // Log and terminate if there is an error preparing the statement
	}

	// Execute the prepared statement
	_, err = statement.Exec()
	if err != nil {
		log.Fatal(err) // Log and terminate if there is an error executing the statement
	}
}

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func CheckPassword(hashedPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}
