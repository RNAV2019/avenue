package model

func CreateLink(link Link) (*Link, error) {
	err := db.Create(&link).Error
	if err != nil {
		return nil, err
	}

	return &link, nil
}
