package course

import (
	"context"
	"io"
	"studyhub/internal/config"
	dto "studyhub/internal/dto/course"
	repo "studyhub/internal/repository/course"
)

type CourseService interface{
	StoreFile(ctx context.Context, file io.ReadSeeker ,req dto.UploadFileRequest)(int, int, error)
    ShareCourse(ctx context.Context, req *dto.ShareCourseRequest, userID int)(int, error)
	UpdateCourse(ctx context.Context, req dto.UpdateCourseRequest, id int, userID int)(int, error)
	DeleteCourse(ctx context.Context, id int, userID int)(int, error)
	ShowAllCourses(ctx context.Context)(int, []dto.ShowCoursesResponse ,error)
	GetURL(ctx context.Context, objectKey string, baseUrl string) (string, error)
	Upload(ctx context.Context, objectKey string,file io.Reader, baseUrl string) error
}

type courseService struct{
	cfg *config.Config
	repo repo.CourseRepository
}

func NewCourseService(cfg *config.Config, repo repo.CourseRepository) CourseService{
	return &courseService{
		cfg: cfg,
		repo: repo,
	}
}