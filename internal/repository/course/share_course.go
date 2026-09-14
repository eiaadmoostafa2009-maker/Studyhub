package course

import (
	"context"
	"errors"
	"studyhub/internal/models"
)

func (r *courseRepository) ShareCourse(ctx context.Context, course models.Course)(error){
	err := r.db.WithContext(ctx).Create(&course).Error

	if err != nil{
		return errors.New("failed to save course in database")
	}

	return nil
}