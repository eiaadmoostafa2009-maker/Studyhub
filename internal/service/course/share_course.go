package course

import (
	"context"
	"net/http"
	dto "studyhub/internal/dto/course"
	"studyhub/internal/models"
)

func (s *courseService) ShareCourse(ctx context.Context, req *dto.ShareCourseRequest, userID int)(int, error){
	//store data
    err :=s.repo.ShareCourse(ctx, models.Course{
		UserID: userID,
		Title: req.Title,
		FileID: req.FileID,
	})

	if err != nil{
		return http.StatusInternalServerError, err
	}
	//return 
	return http.StatusCreated, nil
}