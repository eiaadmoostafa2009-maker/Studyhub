package course

import(
	"context"
	"errors"
	"studyhub/internal/models"
	"gorm.io/gorm"
)

func (r *courseRepository) GetSubmissionByID(ctx context.Context, submissionID int) (*models.CourseSubmission, error) {
	var submission models.CourseSubmission

	err := r.db.WithContext(ctx).First(&submission, submissionID).Error
	if err == gorm.ErrRecordNotFound {
		return nil, errors.New("course submission not found")
	}
	if err != nil {
		return nil, errors.New("failed to retrieve course submission from database")
	}
	return &submission, nil
}