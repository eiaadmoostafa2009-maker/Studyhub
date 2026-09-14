package user

import (
	"context"
	"errors"
	"studyhub/internal/models"
)

func (r *userRepository) DeleteRefreshTokenByUserID(ctx context.Context, userID int) error {
	var token models.RefreshToken
	result := r.db.Where("user_id",userID).Delete(&token)
	if err :=result.Error;err != nil {
		return errors.New("failed to delete due to internal server error")
	}


	return nil
}