package user

import (
	"context"
	"studyhub/internal/config"
	dto "studyhub/internal/dto/user"
	"studyhub/internal/models"
	"studyhub/internal/repository/user"
)

type UserService interface {
    Register(ctx context.Context, req *dto.RegisterRequest) (int, error)
	Login(ctx context.Context, req *dto.LoginRequest) (string, string, int, error)
	SeedAdmin() (int, error)
	GetUserByID(ctx context.Context, id int) (int, *models.User, error)
	DeleteUser(ctx context.Context, id int) (int, error)
	RefreshToken(ctx context.Context, req *dto.RefreshTokenRequest, userID int, refreshToken string) (string, string, int, error)
}

type userService struct {
	config *config.Config
	repo user.UserRepository
}

func NewUserService(config *config.Config, repo user.UserRepository) UserService {
	return &userService{config: config, repo: repo}
}