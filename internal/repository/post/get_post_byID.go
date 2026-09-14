package post

import(
	"context"
	"studyhub/internal/models"
)
func (r *postRepository) GetPostByID(ctx context.Context,postID int, userID int) (*models.Post, error){
	//define the query
	var post models.Post

	err := r.db.WithContext(ctx).Preload("Members").Where("user_id = ? and id = ? ", userID,postID).First(&post).Error
    if err != nil{
		return nil, err
	}

	return &post, nil

}