package user

import (
	"context"
	"errors"
	"studyhub/internal/models"
)

func(r *userRepository) DeleteByName(ctx context.Context, name string) error{
	var user models.User
	result := r.db.WithContext(ctx).Where("name=?", name).Delete(&user)
	if err := result.Error; err != nil {
		return errors.New("failed to delete due tp internal server error")
	}

	if result.RowsAffected == 0 {
		return errors.New("failed to delete user")
	}

	return nil
} 