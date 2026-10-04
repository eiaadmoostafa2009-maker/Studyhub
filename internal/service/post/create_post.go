package post

import(
	"context"
	"studyhub/internal/models"
	"time"
	dto "studyhub/internal/dto/post"
	"net/http"
)
func (s *postService) CreatePost(ctx context.Context, req *dto.CreatePostRequest,userID int, userName string) (int, error){
	//store data
	err :=s.postRepo.StorePost(ctx,&models.Post{
		UserID: userID,
		Title: req.Title,
		Content: req.Content,
		CreatedAt: time.Now(),
		Name: userName,
	})

	if err != nil{
		return http.StatusInternalServerError, err
	}

	//return
	return http.StatusCreated, nil
}