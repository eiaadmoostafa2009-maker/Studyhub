package user

import (
	"context"
	"net/http"
	"studyhub/internal/models"
)

func (s *userService) GetUserByID(ctx context.Context, id int) (int, *models.User, error) {
	user, err := s.repo.GetUserByID(ctx, id)
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	return http.StatusOK, user, nil
}