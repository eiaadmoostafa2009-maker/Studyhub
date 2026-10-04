package post

import(
	"context"
	"errors"
	"studyhub/internal/models"
)
func (r *postRepository) UpdatePost(ctx context.Context, userID int, postID int,post *models.Post) error{
	result := r.db.WithContext(ctx).
    Model(&models.Post{}).
    Where("id = ?", postID, userID).
    Updates(map[string]interface{}{
        "title":   post.Title,
        "content": post.Content,
    })
    
	rowAffected := result.RowsAffected
    if rowAffected == 0{
		return errors.New("failed to update")
	}

	return nil
	
}