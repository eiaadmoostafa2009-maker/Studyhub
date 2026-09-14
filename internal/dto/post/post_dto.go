package post

import (
	"studyhub/internal/models"
	"time"
)
type(
	CreatePostRequest struct {
		Title       string `json:"title" form:"title" validate:"required"`
		Content     string `json:"content" form:"content" validate:"required"`
	}
	
	//update
	UpdatePostRequest struct {
		Title       string `json:"title" form:"title" validate:"required"`
		Content     string `json:"content" form:"content" validate:"required"`
	}
	
	//details
	PostDetailsResponse struct {
		ID int `json:"id"`
		UserID int `json:"user_id"`
		Title string `json:"title"`
		Content string `json:"content"`
		Members int `json:"member"`
		CreatedAt time.Time `json:"created_at"`
	}


	ShowPostDTO struct {
       ID           uint
       Title        string
       Content      string
       UserName   string
       ReleaseDate  time.Time
       Members []models.User
       IsJoined     bool
    }
)