package car

import (
	"github.com/diogocarasco/golang-api-template-v2/internal/db"
	"gorm.io/gorm"
)

func Delete(id string) (rows int64, err error) {

	db, err := db.OpenConnection()
	if err != nil {
		return rows, err
	}

	db.Transaction(func(tx *gorm.DB) error {

		result := tx.Where("id=?", id).Delete(&id)
		if result.Error != nil {
			return result.Error
		}

		rows = tx.RowsAffected

		return nil

	})

	return rows, err

}
