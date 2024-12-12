package database

import (
	"fmt"
	"log"
)

func UpdateShop(user_id, bike_id, total int) error {
	db := SetupDatabase()
	defer db.Close()

	// Requête de mise à jour
	query := "UPDATE shop SET total = ? WHERE user_id = ? AND bike_id = ?"

	// Préparation de la requête
	stmt, err := db.Prepare(query)
	if err != nil {
		return fmt.Errorf("erreur lors de la préparation de la requête : %v", err)
	}
	defer stmt.Close()

	// Exécution de la requête avec les valeurs des paramètres
	_, err = stmt.Exec(total, user_id, bike_id)
	if err != nil {
		return fmt.Errorf("erreur lors de l'exécution de la requête : %v", err)
	}

	fmt.Println("Mise à jour réussie")
	return nil
}

func UpdateStatus(id int, status string)error{
	db := SetupDatabase()
	defer db.Close()

	query := " UPDATE bikes SET status = ? WHERE id = ?"

	stmt, err := db.Prepare(query)
	if err != nil {
		return fmt.Errorf("erreur lors de la préparation de la requête Update status : %v", err)
	}
	defer stmt.Close()
	_, err = stmt.Exec(status, id)
	if err != nil{
		return fmt.Errorf("erreur lors de l'exécution de la requête : %v", err)
	}
	log.Println(id, status)
	return nil
}