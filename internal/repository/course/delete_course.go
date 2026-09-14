package course

import (
	"context"
	"errors"
	"studyhub/internal/models"
)

func(r *courseRepository) DeleteCourse(ctx context.Context, id int, userID int) error{
	result := r.db.WithContext(ctx).Where("user_id=?", userID).Delete(&models.Course{}, id)
   
	if err := result.Error; err != nil{
		return errors.New("failed to update file in the repository")
	}

	rowsAffected := result.RowsAffected

	if rowsAffected == 0{
		return errors.New("failed to delete")
	}

	return nil
}