package course

import (
	"context"
	"errors"
	"studyhub/internal/models"
)

func (r *courseRepository) GetCourseByID(ctx context.Context, id int)(*models.Course, error){
	var course models.Course

	err := r.db.WithContext(ctx).Where("id =?", id).First(&course).Error
	if err != nil{
		return nil, errors.New("failed to find the course in database")
	}

	if course.ID == 0{
		return nil, errors.New("course not found")
	}

	return &course, nil
}