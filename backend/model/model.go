package model

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

type User struct {
	gorm.Model
	FirebaseID string `gorm:"unique;not null"`
	Avenue     Avenue `json:"avenue" gorm:"foreignKey:UserID"`
}

type Avenue struct {
	gorm.Model
	Description string `json:"description" gorm:"not null"`
	UserID      uint   `json:"user_id"`
	Links       []Link `json:"links" gorm:"foreignKey:AvenueID"`
}

type Link struct {
	gorm.Model
	URL         string `json:"url" gorm:"not null"`
	Description string `json:"description" gormå:"not null"`
	AvenueID    uint   `json:"avenue_id"`
}

func Setup() {
	dsn := "postgres://admin:test@localhost:5432/admin?sslmode=disable"
	// dsn := "postgres://admin:test@" + os.Getenv("DB_HOST") + ":5432/admin?sslmode=disable"
	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	err = db.AutoMigrate(&User{}, &Avenue{}, &Link{})
	if err != nil {
		fmt.Println(err)
	}
}
