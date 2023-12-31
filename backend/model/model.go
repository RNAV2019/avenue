package model

import (
	"fmt"
	"os"

	"github.com/google/uuid"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

type Avenue struct {
	gorm.Model
	Description  string      `json:"description" gorm:"not null"`
	ProfileImage string      `json:"profile_image" gorm:"not null"`
	Name         string      `json:"name" gorm:"not null"`
	UserID       uuid.UUID   `json:"user_id" gorm:"type:uuid;not null"`
	Links        []Link      `json:"links" gorm:"foreignKey:AvenueID"`
	Statistics   []Statistic `json:"statistics" gorm:"foreignKey:AvenueID"`
}

type Statistic struct {
	gorm.Model
	GeographicLocation string `json:"geographic_location"`
	TrafficSource      string `json:"traffic_source"`
	ClickTimestamp     string `json:"click_timestamp"`
	AvenueID           uint   `json:"avenue_id"`
}

type Link struct {
	gorm.Model
	URL         string `json:"url" gorm:"not null"`
	Description string `json:"description" gorm:"not null"`
	AvenueID    uint   `json:"avenue_id"`
}

func Setup() {
	dsn := "user=postgres password=" + os.Getenv("DB_PASS") + " host=" + os.Getenv("DB_HOST") + " port=5432 dbname=postgres"
	fmt.Println(dsn)
	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	err = db.AutoMigrate(&Avenue{}, &Statistic{}, &Link{})
	if err != nil {
		panic(err)
	}

	// Check if the foreign key constraint exists before attempting to add it.
	var constraintExists bool
	result := db.Raw(`SELECT EXISTS (
		SELECT 1
		FROM information_schema.table_constraints
		WHERE constraint_name = 'avenues_user_id_fkey'
		AND table_name = 'avenues'
		AND constraint_type = 'FOREIGN KEY'
	)`).Scan(&constraintExists)

	if result.Error != nil {
		return
	}

	// If the foreign key constraint does not exist, execute the ALTER TABLE statement.
	if !constraintExists {
		addForeignKeySQL := `
		ALTER TABLE avenues
		ADD CONSTRAINT avenues_user_id_fkey
		FOREIGN KEY (user_id)
		REFERENCES auth.users (id)
	`

		err = db.Exec(addForeignKeySQL).Error
		if err != nil {
			return
		}

		fmt.Println("Foreign key added successfully!")
	} else {
		fmt.Println("Foreign key constraint already exists.")
	}
}
