package internalsql

import (
	"fmt"

	"studyhub/internal/config"
	"studyhub/internal/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func DBConnect(cfg *config.Config) (*gorm.DB, error) {
	// WAL improves concurrent read behavior. busy_timeout lets SQLite wait briefly
	// instead of immediately returning "database is locked" during contention.
	dsn := cfg.DatabasePath + "?_busy_timeout=5000&_journal_mode=WAL&_foreign_keys=on"

	db, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("open database: %w", err)
	}

	fmt.Println("Connected successfully!")

	err = db.AutoMigrate(
		&models.User{},
		&models.Student{},
		&models.Teacher{},
		&models.Course{},
		&models.CourseSubmission{},
		&models.RefreshToken{},
		&models.Post{},
	)
	if err != nil {
		return nil, fmt.Errorf("migration failed: %w", err)
	}

	return db, nil
}
