package course

import (
	"context"
	"errors"
	"studyhub/internal/models"
)

func(r *courseRepository) ShowAllCourses(ctx context.Context)([]models.Course,error){
	var courses []models.Course

	err := r.db.WithContext(ctx).
		Table("courses").
		Select(`
			courses.id,
			courses.user_id,
			courses.title,
			courses.created_at,
			courses.youtube_url,
			courses.youtube_id,
			users.name as user_name
		`).
		Joins("LEFT JOIN users ON users.id = courses.user_id").
		Order("courses.created_at DESC").
		Scan(&courses).Error

	if err != nil{
		return nil, errors.New("failed to get feed from our database")
	}

	return courses, nil
}