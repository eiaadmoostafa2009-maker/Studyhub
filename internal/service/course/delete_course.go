package course

import (
	"context"
	"errors"
	"net/http"
)

func (s *courseService) DeleteCourse(ctx context.Context, id int, userID int)(int, error){
	//chech if course exists
    course,err := s.repo.GetCourseByID(ctx, id)
	if err != nil{
		return http.StatusInternalServerError, err
	}

	if course.ID == 0{
		return http.StatusBadRequest, errors.New("course not found")
	}
	//delete course
    err = s.repo.DeleteCourse(ctx, id, userID)
	if err != nil{
		return http.StatusInternalServerError,err
	}
	//return
	return http.StatusNoContent,nil
}