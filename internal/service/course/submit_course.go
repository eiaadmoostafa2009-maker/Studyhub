package course

import (
	"context"
	"fmt"
	"io"
	"net/http"

	dto "studyhub/internal/dto/course"
	"studyhub/internal/models"
)

func (s *courseService) SubmitCourse(ctx context.Context, req dto.SubmitCourseRequest, file io.Reader, filename string, userID int) (int, error) {
	submission := models.CourseSubmission{
		UserID: userID,
		Title:     req.Title,
		Status:    models.StatusPending,
	}

	if err := s.repo.CreateSubmission(ctx, &submission); err != nil {
		return http.StatusInternalServerError, err
	}

	caption := fmt.Sprintf(
		"New course submission\nUser ID: %d\nTitle: %s\nSubmission ID: %d",
		submission.UserID,
		submission.Title,
		submission.ID,
	)

	if err := s.telegram.SendVideo(
		ctx,
		file,
		filename,
		caption,
	); err != nil {
		return http.StatusInternalServerError, err
	}

	


	return http.StatusCreated, nil
}