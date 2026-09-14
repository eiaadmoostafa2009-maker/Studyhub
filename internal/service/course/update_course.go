package course

import (
	"context"
	"errors"
	"net/http"
	"os"
	dto "studyhub/internal/dto/course"
	"studyhub/internal/models"
)

func(s *courseService) UpdateCourse(ctx context.Context, req dto.UpdateCourseRequest, id int, userID int)(int, error){
	//check if course exists
    course, err :=s.repo.GetCourseByID(ctx, id)
	if err != nil{
		return http.StatusInternalServerError, err
	}

	user, err := s.repo.GetUserByID(ctx, userID)
	if err != nil{
		return http.StatusInternalServerError, err
	}

	//check if user ids are the same
    if userID != course.UserID || user != os.Getenv("ADMIN_NAME"){
		return http.StatusBadRequest,errors.New("failed to find the course")
	}
	//update data
    err =s.repo.UpdateCourse(ctx, id, &models.Course{
         Title: req.Title,
         FileID: req.FileID,
	})
	if err != nil {
       return http.StatusInternalServerError, err
    }

    return http.StatusOK, nil
}