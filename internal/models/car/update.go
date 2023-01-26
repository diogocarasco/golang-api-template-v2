package car

import (
	"github.com/diogocarasco/golang-api-template-v2/internal/db"
	"gorm.io/gorm"
)

func Update(car Car) (rows int64, err error) {

	db, err := db.OpenConnection()
	if err != nil {
		return rows, err
	}

	db.Transaction(func(tx *gorm.DB) error {

		res := tx.Save(&car)
		if res.Error != nil {
			return res.Error
		}

		rows = tx.RowsAffected

		return nil

	})

	return rows, err

}
