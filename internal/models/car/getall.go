package car

import "github.com/diogocarasco/golang-api-template-v2/internal/db"

func GetAll() (cars []Car, err error) {

	db, err := db.OpenConnection()
	if err != nil {
		return nil, err
	}

	db.Select("model", "brand", "year").Find(&cars)
	if err != nil {
		return cars, err
	}

	return cars, err

}
