package database

import (
	"database/sql"
	"log"
)

func GetUserId(username string)int{
	db := SetupDatabase()
	defer db.Close()

	query := "SELECT user_id FROM users WHERE username = ?"
	stmt, err := db.Prepare(query)
	if err != nil{
		log.Println("Erreur lors de la préparation de la requête:", err)
		return 0
	}
	defer stmt.Close()
	var id int
	err = stmt.QueryRow(username).Scan(&id)
	if err != nil {
		if err == sql.ErrNoRows {
			return 0
		} else {
			log.Println("Erreur lors de l'exécution de la requête:", err)
			return 0
		}
	}
	return id
}

func CheckUser(username, password string) bool {
	db := SetupDatabase()
	defer db.Close()

	// Préparer la requête pour éviter les injections SQL
	query := "SELECT password FROM users WHERE username = ?"
	stmt, err := db.Prepare(query)
	if err != nil {
		log.Println("Erreur lors de la préparation de la requête:", err)
		return false
	}
	defer stmt.Close()

	var hashedPassword string
	// put into the variable the hashedpassword from the db thanks to the users name
	err = stmt.QueryRow(username).Scan(&hashedPassword)
	if err != nil {
		if err == sql.ErrNoRows {
			return false
		} else {
			log.Println("Erreur lors de l'exécution de la requête:", err)
			return false
		}
	}

	// Comparer le mot de passe donné avec le hachage stocké
	return CheckPassword(hashedPassword, password)
}
