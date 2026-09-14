package post

import (
	"context"
	"studyhub/internal/config"
	dto "studyhub/internal/dto/post"
	"studyhub/internal/models"
	repo "studyhub/internal/repository/post"
)

type PostService interface {
	CreatePost(ctx context.Context, post *dto.CreatePostRequest, userID int) (int, error)
	UpdatePost(ctx context.Context, post *dto.UpdatePostRequest, userID, postID int) (int, error)
	DeletePost(ctx context.Context, userID int, postID int) (int, error)
	JoinPost(ctx context.Context, userID, postID int) (int, error)
	DetailPost(ctx context.Context, postID int, userID int) (*models.Post, int, error)
	ShowAllPosts(ctx context.Context, userID int) (int, []dto.ShowPostDTO,error)
}

type postService struct {
	cfg      *config.Config
	postRepo repo.PostRepository
}

func NewPostService(cfg *config.Config, postRepo repo.PostRepository) PostService {
	return &postService{
		cfg:      cfg,
		postRepo: postRepo,
	}
}

