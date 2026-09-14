package user

import (
	"context"
	"errors"
	"net/http"
	dto "studyhub/internal/dto/user"
	"studyhub/internal/models"
	"studyhub/pkg/jwt"
	"time"
	"github.com/google/uuid"

	"golang.org/x/crypto/bcrypt"
)

func (s *userService) Login(ctx context.Context, req *dto.LoginRequest) (string, string, int, error) {
	//check if user exists
	user, err := s.repo.GetUserByEmail(ctx, req.Email)
	if user == nil {
    return "", "", http.StatusUnauthorized,
        errors.New("invalid credentials")
    }

	if err != nil {
		return "", "", 404, err
	}
	//check if password is correct
	bcryptErr := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if bcryptErr != nil {
		return "", "", 401, errors.New("invalid password")
	}
	//generate tokens
	token, err := jwt.GenerateToken(user.ID, user.Name, s.config.SecretKey)
	if err != nil {
		return "", "", 500, err
	}
	refreshToken, err := jwt.GenerateRefreshToken(user.ID, s.config.SecretKey)
	if err != nil {
		return "", "", 500, err
	}

	//get existing refresh token
	existingRefreshToken, err := s.repo.GetRefreshTokenByUserID(ctx, user.ID)
	
	if err != nil {
		return "", "", 500, err
	}

	//generate new refresh token if existing one is expired or not found
	if existingRefreshToken == nil || existingRefreshToken.ExpiresAt.Before(time.Now()) {
		refreshToken, err = jwt.GenerateRefreshToken(user.ID, s.config.SecretKey)
		if err != nil {
			return "", "", 500, err
		}
	}
    
	//store refresh token in database
	_, err = s.repo.StoreRefreshToken(ctx, &models.RefreshToken{
		ID:           uuid.New().String(),
		RefreshToken:        refreshToken,
		UserID:    user.ID,
		ExpiresAt: time.Now().Add(7*24*time.Hour), // 7 days
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})
	if err != nil {
		return "", "", 500, err
	}
	return token, refreshToken, 200, nil
}