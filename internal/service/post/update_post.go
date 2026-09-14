package post

import (
	"context"
	"errors"
	"net/http"
	dto "studyhub/internal/dto/post"
	"studyhub/internal/models"
)
func (s *postService) UpdatePost(ctx context.Context,req *dto.UpdatePostRequest, postID,userID int)(int, error){
	//check if post exists
    post, err := s.postRepo.GetPostByID(ctx, postID, userID)

    if err != nil{
		return http.StatusInternalServerError, err
	}

	if post.ID == 0{
		return http.StatusNotFound, errors.New("post not found")
	}

	if post.UserID != userID{
		return http.StatusNotFound, errors.New("post not found")
	}
	//update post
    err =s.postRepo.UpdatePost(ctx, userID, postID,&models.Post{
		Title: req.Title,
		Content: req.Content,
	})
	if err != nil{
		return http.StatusInternalServerError,err
	}
	//return values
	return http.StatusAccepted,nil
}