package user

import (
	"context"
	"errors"
	"studyhub/internal/models"
)

func (r *userRepository) CreateUser(ctx context.Context, user *models.User) error {
	//query to insert user into database
	result :=r.db.WithContext(ctx).Create(&user)
	if result.Error != nil {
		return  errors.New("failed to create user due to internal server error")
	}
    
	return nil
}
