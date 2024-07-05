package helper

import "gorm.io/gorm"

// สำหรับ Transaction
func Transaction(db *gorm.DB, handler func(*gorm.DB) error) error {
	tx := db.Begin()
	if tx.Error != nil {
		return tx.Error
	}

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			panic(r) // ต่อให้มี panic ก็ตาม, ต้อง Rollback ให้เรียบร้อย
		} else if err := tx.Commit().Error; err != nil {
			tx.Rollback()
		}
	}()

	if err := handler(tx); err != nil {
		tx.Rollback()
		return err
	}

	return nil
}
