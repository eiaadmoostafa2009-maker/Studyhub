package post 

import (
	"context"
	"net/http"
)
func (s *postService) JoinPost (ctx context.Context, userID, postID int)(int, error){
    //check if post exists
    _, err :=s.postRepo.GetPostByID(ctx, postID, userID)
	if err != nil{
		return http.StatusInternalServerError, err
	}
   //check if user already liked it
    isJoined, err :=s.postRepo.IsJoined(ctx, userID, postID)
	if err != nil{
		return http.StatusInternalServerError, err
	}
	
	if isJoined {
    err := s.postRepo.DeleteJoinPost(ctx, userID, postID)
    if err != nil {
        return http.StatusInternalServerError, err
    }

    return http.StatusOK, nil
    }

   err = s.postRepo.JoinPost(ctx, userID, postID)
   if err != nil {
      return http.StatusInternalServerError, err
    }
	
	
	//return
	return http.StatusAccepted, nil
}