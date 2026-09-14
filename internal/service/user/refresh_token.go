package user

import (
	"context"
	"errors"
	"net/http"
	dto "studyhub/internal/dto/user"
	"studyhub/internal/models"
	"studyhub/pkg/jwt"
	"time"
)

func (s *userService) RefreshToken(ctx context.Context, req *dto.RefreshTokenRequest, userID int, refreshToken string) (string, string, int, error) {

    //check if user exists
	user , err := s.repo.GetUserByID(ctx, userID)
	if err != nil {
		return "", "", 0, err
	}
	
	if user == nil {
		return "", "", 0, errors.New("user not found")
	}
	//get refresh token by user id
    refreshTokenExists, err := s.repo.GetRefreshTokenByUserID(ctx, userID)
	if err != nil {
		return "", "", 0, err
	}

	if refreshTokenExists == nil {
		return "", "", 0, errors.New("refresh token not found or expired")
	}
	//compare the refresh token from request and the one in database
    if refreshTokenExists.RefreshToken != refreshToken {
		return "", "", 0, errors.New("invalid refresh token")
	}

	if refreshTokenExists.ExpiresAt.Before(time.Now()) {
        return "", "", http.StatusUnauthorized,
        errors.New("refresh token expired")
    }

	token,err := jwt.GenerateToken(userID, user.Name, s.config.SecretKey)
	if err != nil {
		return "", "", 0, err
	}
	//if they match, generate new token and refresh token
    refreshToken,err = jwt.GenerateRefreshToken(userID, s.config.SecretKey)
	if err != nil {
		return "", "", 0, err
	}
	//delete old tokens
	err = s.repo.DeleteRefreshTokenByUserID(ctx, userID)

	//store new refresh token in database
	now := time.Now()
	RefreshToken := &models.RefreshToken{
		UserID:       userID,
		RefreshToken: refreshToken,
		ExpiresAt:    now.Add(7 * 24 * time.Hour),
		CreatedAt:    now,
		UpdatedAt:    now,
	}
	_,err = s.repo.StoreRefreshToken(ctx, RefreshToken)
	if err != nil {
		return "", "", 0, err
	}
	return refreshToken, token, 201, nil
}