package database

import (
	"database/sql"
	"fmt"
	"log"
)

type Bike struct {
	ID        int
	ImagePath string
	BikeType  string
	MotorType string
	Status    string
	Size      float64
	Speed     int
	Autonomy  int
	Price     float64
	Battery   int
}

func GetUserId(username string) int {
	db := SetupDatabase()
	defer db.Close()

	query := "SELECT user_id FROM users WHERE username = ?"
	stmt, err := db.Prepare(query)
	if err != nil {
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

// Fonction pour récupérer tous les vélos de la base de données
func GetAllBikes() ([]Bike, error) {
	db := SetupDatabase()
	defer db.Close()

	// Requête SQL pour récupérer tous les vélos
	query := "SELECT id, image_path, bike_type, motor_type, size, speed, autonomy, battery, price, status FROM bikes"

	// Exécuter la requête
	rows, err := db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("erreur lors de l'exécution de la requête : %v", err)
	}
	defer rows.Close()
	// Slice pour stocker les vélos
	var bikes []Bike

	// Itérer sur les lignes retournées par la requête
	for rows.Next() {
		var bike Bike
		err := rows.Scan(
    &bike.ID,
    &bike.ImagePath,
    &bike.BikeType,
    &bike.MotorType,
    &bike.Size,
    &bike.Speed,
    &bike.Autonomy,
    &bike.Battery,
    &bike.Price,
    &bike.Status,
)

		if err != nil {
			return nil, fmt.Errorf("erreur lors de l'analyse des données : %v", err)
		}

		// Ajouter le vélo au slice
		bikes = append(bikes, bike)
	}

	// Vérifier les erreurs lors de l'itération
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("erreur pendant l'itération : %v", err)
	}
	return bikes, nil
}

func GetUserRole(username string) string {
	db := SetupDatabase()
	defer db.Close()

	query := "SELECT role FROM users WHERE username = ?"
	stmt, err := db.Prepare(query)
	if err != nil {
		log.Println("Erreur lors de la préparation de la requête:", err)
		return ""
	}
	defer stmt.Close()
	var role string
	err = stmt.QueryRow(username).Scan(&role)
	if err != nil {
		if err == sql.ErrNoRows {
			return ""
		} else {
			log.Println("Erreur lors de l'exécution de la requête:", err)
			return ""
		}
	}
	return role
}