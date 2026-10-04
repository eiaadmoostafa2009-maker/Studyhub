package course

import (
	"context"
	"net/http"
	dto "studyhub/internal/dto/course"

)

func(s *courseService) ShowAllCourses(ctx context.Context)(int, []dto.ShowCoursesResponse,error){
	courses, err := s.repo.ShowAllCourses(ctx)
    if err != nil {
        return http.StatusInternalServerError, nil, err
    }

    result := make(
        []dto.ShowCoursesResponse,
        0,
        len(courses),
    )

    for _, course := range courses {
        user, err := s.repo.GetUserByID(ctx, course.UserID)
        if err != nil{
            return http.StatusInternalServerError, nil, err
        }
        
        result = append(result, dto.ShowCoursesResponse{
             ID:          course.ID,
             Title:       course.Title,
             UserID:   course.UserID,
             UserName: user.Name,
             YoutubeUrl:  course.YoutubeUrl,
        })
        
            
        
    }

    return http.StatusOK, result, nil
}