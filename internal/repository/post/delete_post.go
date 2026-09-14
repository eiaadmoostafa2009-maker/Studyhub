package post

import (
	"context"
	"errors"
	"studyhub/internal/models"

)

func (r *postRepository) DeletePost(ctx context.Context,userID, postID int) error{
	var post models.Post
	result := r.db.WithContext(ctx).Where("id = ? AND (user_id = ? OR ? IN (SELECT id FROM users WHERE role = 'admin'))", postID, userID, userID).Delete(&post)
	if err := result.Error; err != nil{
		return err
	}
    
	if result.RowsAffected == 0{
		return errors.New("didn't find the post to delete it")
	}
	return nil
}