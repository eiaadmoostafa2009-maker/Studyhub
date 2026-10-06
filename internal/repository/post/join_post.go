package post

import (
	"context"
	"studyhub/internal/models"
	"errors"
)

func (r *postRepository) JoinPost(ctx context.Context, userID, postID int) error {
	var post models.Post
    if err := r.db.WithContext(ctx).First(&post, postID).Error; err != nil {
        return err
    }

    var user models.User
    if err := r.db.WithContext(ctx).First(&user, userID).Error; err != nil {
        return err
    }

    err:= r.db.WithContext(ctx).Model(&post).Association("Members").Append(&user)
	if err != nil{
		return errors.New("our db failed to sign you as a joiner in this post, sorry")
	}

	return nil
}