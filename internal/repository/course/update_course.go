package course

import (
	"context"
	"errors"
	"studyhub/internal/models"
)

func (r *courseRepository) UpdateCourse(ctx context.Context, id int, course *models.Course) error{
	result :=r.db.WithContext(ctx).Model(&models.Course{}).Where("id=?",id).Updates(course)
    if err:=result.Error;err !=nil{
		return err
	}

	rowsAffected :=result.RowsAffected
	if rowsAffected == 0{
		return errors.New("failed to update")
	}
	return nil
}