package post

import (
	"context"
	"net/http"
	dto "studyhub/internal/dto/post"
)

func (s *postService) ShowAllPosts(ctx context.Context, userID int)(int, []dto.ShowPostDTO, error){
	 //get all posts
	 posts ,err := s.postRepo.ShowAllPosts(ctx, userID)
	 if err != nil{
		return http.StatusInternalServerError, nil, err
	 }
	  result := make([]dto.ShowPostDTO, 0, len(posts))

    for _, post := range posts {

        result = append(result, dto.ShowPostDTO{
            ID:           uint(post.ID),
            Title:        post.Title,
            Content:      post.Content,
            UserName:   post.UserName,
            ReleaseDate:  post.CreatedAt,
            Members: post.Members,
            IsJoined:     post.IsJoined,
        })
    }

	 //return values
	 return http.StatusOK, result, nil
}