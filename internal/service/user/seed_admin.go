package user

import "net/http"

func (s *userService) SeedAdmin() (int, error) {
	err := s.repo.SeedAdmin()
	if err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusCreated, nil
}
