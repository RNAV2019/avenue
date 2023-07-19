package model

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

// type User struct {
// 	gorm.Model
// 	ID           uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
// 	Email        string    `gorm:"uniqueIndex"`
// 	PasswordHash string
// 	Avenue       Avenue `json:"avenue" gorm:"foreignKey:UserID"`
// }

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
	GeographicLocation string    `json:"geographic_location"`
	TrafficSource      string    `json:"traffic_source"`
	ClickTimestamp     time.Time `json:"click_timestamp"`
	AvenueID           uint      `json:"avenue_id"`
}

type Link struct {
	gorm.Model
	URL         string `json:"url" gorm:"not null"`
	Description string `json:"description" gorm:"not null"`
	AvenueID    uint   `json:"avenue_id"`
}

func Setup() {
	dsn := "postgresql://postgres:postgres@localhost:54322/postgres"
	// dsn := "postgres://admin:test@" + os.Getenv("DB_HOST") + ":5432/admin?sslmode=disable"
	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	err = db.AutoMigrate(&Avenue{}, &Statistic{}, &Link{})
	if err != nil {
		fmt.Println(err)
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
		fmt.Println("failed to check foreign key constraint:", result.Error)
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
			fmt.Println("failed to add foreign key:", err)
			return
		}

		fmt.Println("Foreign key added successfully!")
	} else {
		fmt.Println("Foreign key constraint already exists.")
	}
}
