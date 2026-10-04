package course

import (
	"context"
	"errors"
	"studyhub/internal/models"
)

func (r *courseRepository) CreateSubmission(ctx context.Context, submission *models.CourseSubmission) error {
	err := r.db.WithContext(ctx).Create(&submission).Error
	if err != nil {
		return errors.New("failed to create submition in db")
	}
	return nil
}