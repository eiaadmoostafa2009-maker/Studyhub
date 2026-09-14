package post

import(
	"context"
)

func (r *postRepository) IsJoined(ctx context.Context, userID, postID int) (bool, error){

    var count int64

    err := r.db.WithContext(ctx).
        Table("post_members").
        Where("user_id = ? AND post_id = ?", userID, postID).
        Count(&count).Error

    return count > 0, err
}
