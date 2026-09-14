package user

import (
	"context"
	"studyhub/internal/models"

	"gorm.io/gorm"
)

type UserRepository interface {
	GetUserByEmail(ctx context.Context, email string) (*models.User , error)
	CreateUser(ctx context.Context, user *models.User) (error)
    GetRefreshTokenByUserID(ctx context.Context, userID int) (*models.RefreshToken, error)
	StoreRefreshToken(ctx context.Context, refreshToken *models.RefreshToken) (int, error)
	GetUserByID(ctx context.Context, userID int) (*models.User, error)
	DeleteRefreshTokenByUserID(ctx context.Context, userID int) error
	SeedAdmin() error
	DeleteUser(ctx context.Context, id int) error
}

type userRepository struct {
	db *gorm.DB
}

//func to return struct values
func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}
