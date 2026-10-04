package post

import(
	"context"
	"studyhub/internal/models"
)
func (r *postRepository) GetPostByID(ctx context.Context,postID int) (*models.Post, error){
	//define the query
	var post models.Post

	err := r.db.WithContext(ctx).Preload("Members").Where("id = ? ", postID).First(&post).Error
    if err != nil{
		return nil, err
	}

	return &post, nil

}