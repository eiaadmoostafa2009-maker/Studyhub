package course

import (
	"context"
	"net/http"
	dto "studyhub/internal/dto/course"
)

func(s *courseService) ShowAllCourses(ctx context.Context)(int, []dto.ShowCoursesResponse,error){
	//get all posts
	courses,err :=s.repo.ShowAllCourses(ctx)
    if err != nil{
		return http.StatusInternalServerError, nil, err
	}
    
	
	//map the response
	var feed = make([]dto.ShowCoursesResponse, 0, len(courses))
    for _,course :=range courses{
		userName, err := s.repo.GetUserByID(ctx, course.UserID)
		if err != nil{
		    return http.StatusInternalServerError, nil, err
		}

		feed =append(feed, dto.ShowCoursesResponse{
			ID: course.ID,
			UserName: userName,
			Title: course.Title,
			FileURL: course.ObjectKey,
		})
	}
	//return
	
    return http.StatusAccepted, feed, nil
}