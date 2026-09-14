package user

import (
	"context"
	"errors"
	"studyhub/internal/models"

	"gorm.io/gorm"
)

func (r *userRepository) GetUserByID(ctx context.Context, userID int) (*models.User, error) {
	//query to get user by id
	var user models.User
	err := r.db.WithContext(ctx).First(&user, userID).Error
	if err == gorm.ErrRecordNotFound{
		return nil, nil
	}
	if err != nil {
		return nil, errors.New("failed to find the user by id")
	}
	return &user, nil
}