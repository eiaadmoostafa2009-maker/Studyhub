package post

import (
	"context"
	"studyhub/internal/models"
)

func (r *postRepository) DeleteJoinPost(ctx context.Context, userID, postID int) error{
	err := r.db.WithContext(ctx).
    Model(&models.Post{ID: postID}).
    Association("Members").
    Delete(&models.User{ID: userID})
	
	if err != nil{
		return err
	}

    return nil
}