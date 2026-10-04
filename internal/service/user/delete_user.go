package user

import (
	"context"
	"errors"
	"net/http"
)

func (s *userService) DeleteUser(ctx context.Context, id int) (int, error) {
	//check if user exists
	user, err := s.repo.GetUserByID(ctx, id)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	if user == nil {
		return http.StatusBadRequest, errors.New("user doesn't exists")
	}
	//delete user
	err = s.repo.DeleteUser(ctx, id)
	if err != nil {
		return http.StatusInternalServerError, err
	}
	//return values
	return http.StatusOK, nil
}
