package model

func CreateUser(user User) (*User, error) {
	newUser := User{
		FirebaseID: user.FirebaseID,
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

func GetUserByUID(uid string) (*User, error) {
	var user User
	err := db.Preload("Avenue.Links").Where("firebase_id = ?", uid).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func GetAllUsers() ([]User, error) {
	var users []User
	err := db.Preload("Avenue.Links").Find(&users).Error
	if err != nil {
		return nil, err
	}

	return users, nil
}

func DeleteUser(uid string) error {
	return db.Unscoped().Delete(&User{}, uid).Error
}
