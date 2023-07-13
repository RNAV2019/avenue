package model

func CreateAvenue(id uint) error {
	tx := db.Create(&Avenue{
		UserID: id,
	})
	return tx.Error
}
