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
	CreatedAt string
}
type BikeShop struct {
	ID        int
	UserID    string
	BikeID    string
	BikeType  string
	Price	  float64
	Size	  float64
	Total	  int
	CreatedAt string
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

func GetOneBike(id int) ([]Bike, error) {
	db := SetupDatabase()
	defer db.Close()

	// Requête SQL pour récupérer tous les vélos
	query := "SELECT id, image_path, bike_type, motor_type, size, speed, autonomy, battery, price, status FROM bikes WHERE id = ?"
	// Exécuter la requête
	rows, err := db.Query(query, id)
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
func GetShopBike(user_id int) ([]BikeShop, int, error) {
	db := SetupDatabase()
	defer db.Close()

	
	totalAmount := 0         // Somme des totaux de chaque vélo
	query := "SELECT id, user_id, bike_id, bike_type, price, size, total, created_at FROM shop WHERE user_id = ?"
	rows, err := db.Query(query, user_id)
	if err != nil {
		return nil, totalAmount, fmt.Errorf("erreur lors de l'exécution de la requête : %v", err)
	}
	defer rows.Close()
	
	var bikes []BikeShop

	for rows.Next() {
		var bike BikeShop
		err = rows.Scan(
			&bike.ID,
			&bike.UserID,
			&bike.BikeID,
			&bike.BikeType,
			&bike.Price,
			&bike.Size,
			&bike.Total,      // On récupère la colonne `total`
			&bike.CreatedAt,  // On récupère la colonne `created_at`
		)
		if err != nil {
			return nil, totalAmount, fmt.Errorf("erreur lors de l'analyse des données : %v", err)
		}
		bikes = append(bikes, bike) 
		totalAmount += bike.Total    // On additionne la valeur de `total`
	}

	if err := rows.Err(); err != nil {
		return nil, totalAmount, fmt.Errorf("erreur pendant l'itération : %v", err)
	}

	return bikes, totalAmount, nil
}


func VerifBikeId(user_id, bike_id int) (int, bool) {
	db := SetupDatabase()
	defer db.Close() // Defer garantit la fermeture de la base

	// On récupère uniquement le premier bike_id
	query := "SELECT bike_id,total FROM shop WHERE user_id = ? AND bike_id = ?" // On prend le premier bike_id
	var bikeID int
	var total int
	// On exécute la requête avec QueryRow, car on veut seulement 1 résultat
	err := db.QueryRow(query, user_id, bike_id).Scan(&bikeID,
													&total)
	if err != nil {
		// Si on ne trouve rien, on retourne 0 et false
		if err == sql.ErrNoRows {
			return 1, false // Pas de bike_id trouvé
		}
		return 1, false // Erreur SQL
	}
	fmt.Println(total)
	// Si un bike_id est trouvé, on le retourne avec true
	return total, true
}	
