package course

import (
	"context"
	"io"
    "studyhub/internal/integrations/telegram"
	"studyhub/internal/config"
	dto "studyhub/internal/dto/course"
	repo "studyhub/internal/repository/course"
)

type CourseService interface {
    PublishCourse(ctx context.Context, submissionID int, req dto.PublishCourseRequest) (int, error)
	UpdateCourse(ctx context.Context, req dto.UpdateCourseRequest, id int, userID int) (int, error)
	DeleteCourse(ctx context.Context, id int, userID int) (int, error)
	ShowAllCourses(ctx context.Context) (int, []dto.ShowCoursesResponse, error)
	SubmitCourse(ctx context.Context, req dto.SubmitCourseRequest, file io.Reader, filename string, userID int) (int, error)
}

type courseService struct {
	cfg    *config.Config
	repo   repo.CourseRepository
	telegram telegram.Client
}

func NewCourseService(cfg *config.Config, repo repo.CourseRepository, telegram telegram.Client) CourseService {
	return &courseService{
		cfg:    cfg,
		repo:   repo,
		telegram: telegram,
	}
}
