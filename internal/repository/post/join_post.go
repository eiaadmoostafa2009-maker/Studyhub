package post

import (
	"context"
	"studyhub/internal/models"
)

func (r *postRepository) JoinPost(ctx context.Context, userID, postID int) error {
	var err error = r.db.WithContext(ctx).Model(&models.Post{}).Where("id = ?", postID).Association("Members").Append(&models.User{ID: userID})
	if err != nil{
		return err
	}

	return nil
}