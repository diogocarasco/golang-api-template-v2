package car

import (
	"github.com/diogocarasco/golang-api-template-v2/internal/db"
	"gorm.io/gorm"
)

func Create(car Car) (uuid string, err error) {

	db, err := db.GetConnectionPool()
	if err != nil {
		return uuid, err
	}

	/*
		db.Transaction(func(tx *gorm.DB) error {

			res := tx.Create(&car)
			if res.Error != nil {
				return res.Error
			}

			return nil

		})
	*/

	tx, err := db.BeginTx(*gorm.DB)
	res := tx.Create(&car)
	if res.Error != nil {
		return res.Error
	}

	uuid = car.ID

	return uuid, err

}
