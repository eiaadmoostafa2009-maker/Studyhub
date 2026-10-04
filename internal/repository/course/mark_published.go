package course

import (
	"context"
	"studyhub/internal/models"
    "errors"
)

func (r *courseRepository) MarkRequestAsPublished(ctx context.Context, submissionID int, youtubeURL string, youtubeID string) error {
	var submission models.CourseSubmission

	result := r.db.WithContext(ctx).Model(&submission).Where("id = ?", submissionID).Updates(map[string]interface{}{
		"status":      models.StatusApproved,
		"youtube_url": youtubeURL,
	})
	if err := result.Error; err != nil {
		return errors.New("failed to mark submission as published")
	}

	if r.db.RowsAffected == 0 {
		return errors.New("no submission found with the given ID")
	}

	return nil
}