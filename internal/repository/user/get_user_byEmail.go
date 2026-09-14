package user

import (
	"context"
	"errors"
	"studyhub/internal/models"

	"gorm.io/gorm"
)

func (r *userRepository) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	//query to get user by email
	var users models.User
	err := r.db.WithContext(ctx).Where("email", email).First(&users).Error
	if err == gorm.ErrRecordNotFound{
		return nil, nil
	}
	if err != nil {
		return nil, errors.New("failed to find the user by email")
	}
	return &users, nil
}
