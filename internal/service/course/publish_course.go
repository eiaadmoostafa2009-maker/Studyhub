package course

import (
	"context"
	"errors"
	"net/http"
	dto "studyhub/internal/dto/course"
	"studyhub/internal/models"
)

func(s *courseService) PublishCourse(ctx context.Context, submissionID int, req dto.PublishCourseRequest) (int, error){

    submission, err := s.repo.GetSubmissionByID(
        ctx,
        req.SubmissionID,
    )

    if err != nil {
        return http.StatusInternalServerError, err
    }

    if submission.Status != models.StatusPending {
        return http.StatusBadRequest, errors.New("submission is not pending")
    }

    course := models.Course{
        UserID:     submission.UserID,
        Title:      submission.Title,
        YoutubeUrl: req.YoutubeURL,
        YoutubeID:  req.YoutubeID,
    }

    if err := s.repo.ShareCourse(ctx, course); err != nil {
        return http.StatusInternalServerError, err
    }

	err =s.repo.MarkRequestAsPublished(ctx, submission.ID, req.YoutubeURL, req.YoutubeID)
    if err != nil {
        return http.StatusInternalServerError, err
    }

    return http.StatusOK, nil
    
}
