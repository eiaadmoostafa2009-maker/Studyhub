package post

import (
	"context"
	"studyhub/internal/models"
)

func (r *postRepository) ShowAllPosts(ctx context.Context, userID int) ([]models.Post,error) {
    var posts []models.Post

    err := r.db.WithContext(ctx).
        Table("posts").
        Select(`
            posts.id,
            posts.title,
            posts.content,
            posts.created_at,
            users.name AS user_name,
            COUNT(DISTINCT post_members.user_id) AS members_count,
            CASE
                WHEN EXISTS (
                    SELECT 1
                    FROM post_members pm
                    WHERE pm.post_id = posts.id
                    AND pm.user_id = ?
                )
                THEN 1
                ELSE 0
            END AS is_joined
        `, userID).
        Joins("LEFT JOIN users ON users.id = posts.user_id").
        Joins("LEFT JOIN post_members ON post_members.post_id = posts.id").
        Group(`
            posts.id,
            posts.title,
            posts.content,
            posts.created_at,
            users.name
        `).
        Order("posts.created_at DESC").
        Find(&posts).Error

    return posts, err
}

