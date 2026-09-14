package post

import (
	"context"
	"studyhub/internal/models"
)

func (r *postRepository) StorePost(ctx context.Context, post *models.Post) (error){
	result := r.db.WithContext(ctx).Create(&post)
    err := result.Error
	if err != nil{
		return err
	}
    

	return nil
}