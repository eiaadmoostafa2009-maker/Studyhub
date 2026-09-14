package internalsql

import (
	"fmt"
	
	"studyhub/internal/config"
	"studyhub/internal/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func DBConnect(cfg *config.Config) (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("studyhub.db"), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("open database: %w", err)

	}

	// Verify connection
	fmt.Println("Connected successfully!")

	//auto migrate
	err = db.AutoMigrate(
		&models.User{},
		&models.Student{},
		&models.Teacher{},
		&models.Course{},
		&models.File{},
		&models.RefreshToken{},
		&models.Post{},
        )

	if err != nil {
       return nil, fmt.Errorf("migration failed: %w", err)
    }

	return db, nil
}
