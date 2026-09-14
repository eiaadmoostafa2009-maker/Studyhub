package user

import (
	"context"
	"errors"
	"studyhub/internal/models"
)
func (r *userRepository) StoreRefreshToken(ctx context.Context, refreshToken *models.RefreshToken) (int, error) {
	err := r.db.WithContext(ctx).Create(refreshToken).Error
	if err != nil {
		return 0, errors.New("failed to store the refresh token")
	}
	return 200,nil
}