package post

import(
	"context"
	"net/http"
	"errors"
)

func (s *postService) DeletePost(ctx context.Context, userID int, postID int)(int ,error){
	//check if post exists
    post ,err :=s.postRepo.GetPostByID(ctx, postID, userID)
	if err != nil{
		return http.StatusInternalServerError,err
	}
    if post.ID == 0{
       return http.StatusNotFound, errors.New("post not found")
	}
	//delete post
    err = s.postRepo.DeletePost(ctx, userID, postID)
	if err != nil {
		return http.StatusInternalServerError,err
	}
	//return value
	return http.StatusNoContent, nil
}