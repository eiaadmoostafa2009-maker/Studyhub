package user

import (
	"errors"
	"os"
	"studyhub/internal/models"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func (r *userRepository) SeedAdmin() error {
	var count int64

	if err := r.db.Model(&models.User{}).
		Where("role = ?", models.AdminRole).
		Count(&count).Error; err != nil {
		return err
	}

	if count > 0 {
		return nil
	}

	passwordHash, err := bcrypt.GenerateFromPassword(
		[]byte(os.Getenv("ADMIN_PASSWORD")),
		bcrypt.DefaultCost,
	)
	if err != nil {
		return errors.New("failed to create admin account")
	}

	admin := models.User{
	    Name: os.Getenv("ADMIN_NAME"),
		Email:    os.Getenv("ADMIN_EMAIL"),
		Password: string(passwordHash),
		Role:     models.AdminRole,
		CreatedAt: time.Now(),
	}

	return r.db.Create(&admin).Error
}
