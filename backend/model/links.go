package model

func CreateLink(link Link) (*Link, error) {
	err := db.Create(&link).Error
	if err != nil {
		return nil, err
	}

	return &link, nil
}

func DeleteLink(id int) error {
	return db.Unscoped().Delete(&Link{}, id).Error
}

func UpdateLink(link Link) (*Link, error) {
	err := db.Save(&link).Error
	if err != nil {
		return nil, err
	}
	return &link, nil
}
