package post

import (
	"context"
	"studyhub/internal/models"

	"gorm.io/gorm"
)

type PostRepository interface {
	StorePost(ctx context.Context,post *models.Post) (error)
	GetPostByID(ctx context.Context,postID int,userID int) (*models.Post, error)
	UpdatePost(ctx context.Context,userID int,postID int,post *models.Post) error
	DeletePost(ctx context.Context,userID,postID int) error
	IsJoined(ctx context.Context, userID, postID int) (bool, error)
	JoinPost(ctx context.Context,userID, postID int) error
	DeleteJoinPost(ctx context.Context, userID, postID int) error
	ShowAllPosts(ctx context.Context, userID int) ([]models.Post, error)
}

type postRepository struct {
	db *gorm.DB
}

func NewPostRepository(db *gorm.DB) PostRepository {
	return &postRepository{db: db}
}
