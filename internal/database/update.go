package database

import "fmt"

func UpdateShop(user_id, bike_id, total int) error {
	db := SetupDatabase()
	defer db.Close()

	// Requête de mise à jour
	query := "UPDATE shop SET total = ? WHERE user_id = ? AND bike_id = ?"

	// Préparation de la requête
	stmt, err := db.Prepare(query,)
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
