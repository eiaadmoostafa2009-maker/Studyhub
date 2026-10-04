package user

import (
	"context"
	"errors"
	"studyhub/internal/models"
	"gorm.io/gorm"
)

func (r *userRepository) GetRefreshTokenByUserID(ctx context.Context, userID int) (*models.RefreshToken, error) {
	var token models.RefreshToken
	err := r.db.WithContext(ctx).Where("user_id", userID).First(&token).Error
	
	if errors.Is(err, gorm.ErrRecordNotFound) {
       return nil, nil
    }

	if err != nil {
		return nil, errors.New("failed to find the refresh token")
	}
	return &token, nil
}