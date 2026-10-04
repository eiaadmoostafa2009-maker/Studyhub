package user

import (
	"context"
	"database/sql"
	"errors"
	"net/http"
	"studyhub/internal/models"
	"time"
	"golang.org/x/crypto/bcrypt"
	dto "studyhub/internal/dto/user"
)

func (r *userService) Register(ctx context.Context, req *dto.RegisterRequest) (int, error) {
	//check if user already exists
	checkUser, err := r.repo.GetUserByEmail(ctx, req.Email)
	if err != nil && err != sql.ErrNoRows {
		return http.StatusInternalServerError, err
	}
	if checkUser != nil {
		return http.StatusConflict, errors.New("user already exists")
	}
	//hash password
	bytes, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return http.StatusInternalServerError, err
	}
	//create user in database
	user := &models.User{
		Name:   req.UserName,
		Email:      req.Email,
		Password:   string(bytes),
		School: req.School,
		Role:   req.UserRole,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}
	err = r.repo.CreateUser(ctx, user)
	if err != nil {
		return http.StatusInternalServerError, err
	}
	return http.StatusCreated, nil
}
