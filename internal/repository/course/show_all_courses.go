package course

import (
	"context"
	"errors"
	"studyhub/internal/models"
)

func(r *courseRepository) ShowAllCourses(ctx context.Context)([]models.CourseWithFile,error){
	var courses []models.CourseWithFile

	err := r.db.WithContext(ctx).
		Table("courses").
		Select(`
			courses.id,
			courses.user_id,
			courses.title,
			courses.created_at,
			files.object_key
		`).
		Joins("LEFT JOIN users ON users.id = courses.user_id").
		Joins("LEFT JOIN files ON files.id = courses.file_id").
		Order("courses.created_at DESC").
		Scan(&courses).Error

	if err != nil{
		return nil, errors.New("failed to get feed from our database")
	}

	return courses, nil
}