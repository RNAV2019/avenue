package model

func GetCurrentAvenue(user User) Avenue {
	var avenue Avenue
	db.Where("user_id = ?", user.ID).First(&avenue)
	return avenue
}
