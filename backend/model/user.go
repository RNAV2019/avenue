package model

func CreateUser(user *User) (*User, error) {
	newUser := User{
		FBToken: user.FBToken,
		Avenue: Avenue{
			Title:       "Avenue Title",
			Description: "Avenue Description",
		},
	}

	err := db.Create(&newUser).Error
	if err != nil {
		return nil, err
	}

	return &newUser, nil
}

func GetAllUsers() ([]User, error) {
	var users []User
	err := db.Preload("Avenue.Links").Find(&users).Error
	if err != nil {
		return nil, err
	}

	return users, nil
}

func UpdateUser(user User) error {
	tx := db.Save(&user)
	return tx.Error
}
