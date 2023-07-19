package model

import (
	"fmt"

	"github.com/google/uuid"
	supa "github.com/nedpals/supabase-go"
)

func GetAvenueByUID(uid string) (Avenue, error) {
	var avenue Avenue
	tx := db.Preload("Links").Where("user_id = ?", uid).First(&avenue)
	if tx.Error != nil {
		return Avenue{}, tx.Error
	}
	return avenue, nil
}

func GetProfileImageAndName(uid string) (string, string, error) {
	var avenue Avenue
	tx := db.Where("user_id = ?", uid).First(&avenue)
	if tx.Error != nil {
		return "", "", tx.Error
	}
	return avenue.ProfileImage, avenue.Name, nil
}

func CreateAvenue(uid string, user *supa.User) (*Avenue, error) {
	fmt.Println(uuid.MustParse(uid))
	profileURL := user.UserMetadata["avatar_url"].(string)
	name := user.UserMetadata["name"].(string)
	fmt.Println(profileURL)
	fmt.Println(name)

	var avenue Avenue = Avenue{
		UserID:       uuid.MustParse(uid),
		Description:  "New Avenue",
		ProfileImage: profileURL,
		Name:         name,
		Links:        []Link{},
	}

	var avenues []Avenue

	db.Where("user_id = ?", uid).Find(&avenues)
	if len(avenues) == 0 {
		err := db.Create(&avenue).Error
		if err != nil {
			return nil, err
		}
	} else {
		fmt.Println("Avenue already exists")
		return &avenues[0], nil
	}

	return &avenue, nil

}
