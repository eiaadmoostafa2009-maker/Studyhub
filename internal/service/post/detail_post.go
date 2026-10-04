package post

import (
	"context"
	"errors"
	"net/http"
	"studyhub/internal/models"
)

func(s *postService) DetailPost (ctx context.Context, postID int)(*models.Post, int,error){
	//search for post
    post, err :=s.postRepo.GetPostByID(ctx, postID)
	if err != nil{
		return nil, http.StatusInternalServerError, err
	}

	if post == nil{
       return nil, http.StatusInternalServerError, errors.New("post is not found")
	}
	//return the details
	return post, http.StatusOK, nil
}