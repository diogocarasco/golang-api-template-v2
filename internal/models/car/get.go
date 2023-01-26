package car

import (
	"github.com/diogocarasco/golang-api-template-v2/internal/db"
)

func Get(id string) (car Car, err error) {

	db, err := db.OpenConnection()
	if err != nil {
		return car, err
	}

	db.Model(&car).Select("model", "brand", "year").First(&car, `id = $1`, id)
	if err != nil {
		return car, err
	}

	return car, err

}
