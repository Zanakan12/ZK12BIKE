package database

import (
	"log"
	"os"
)

func DeleteBike(bike_id int, fildPath string){
	db := SetupDatabase()
	defer db.Close()

	query := "DELETE FROM bikes WHERE id = ?"

	_,err := db.Exec(query,bike_id)
	if err != nil{
		log.Println("error execute delete bike action on the db")
	}
	fildPath = fildPath[1:]
	err = os.Remove(fildPath)
	if err != nil{
		log.Println("error to delelete file : v%\n", err)
	}
}

func DeleteToCart(bike_id int){
	db := SetupDatabase()
	defer db.Close()

	log.Println("here")
	query := "DELETE FROM shop WHERE bike_id = ?"

	_,err := db.Exec(query,bike_id)
	if err != nil{
		log.Println("error execute delete bike action on the db")
	}

	if err != nil{
		log.Println("error to delelete file : v%\n", err)
	}
}