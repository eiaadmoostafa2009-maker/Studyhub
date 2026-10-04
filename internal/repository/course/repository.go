package course

import (
	"context"
	"studyhub/internal/models"

	"gorm.io/gorm"
)

type CourseRepository interface{
    ShareCourse(ctx context.Context, course models.Course)(error)
	GetCourseByID(ctx context.Context, id int)(*models.Course, error)
	UpdateCourse(ctx context.Context, id int, course *models.Course) error
	DeleteCourse(ctx context.Context, id int, userID int)error
	ShowAllCourses(ctx context.Context)([]models.Course,error)
	GetUserByID(ctx context.Context, userID int) (*models.User, error)
	GetSubmissionByID(ctx context.Context, submissionID int) (*models.CourseSubmission, error)
	MarkRequestAsPublished(ctx context.Context, submissionID int, youtubeURL string, youtubeID string) error
	//RejectRequest(ctx context.Context, submissionID int) error
	CreateSubmission(ctx context.Context, submission *models.CourseSubmission) error
}

type courseRepository struct{
	db *gorm.DB
}

func NewCourseRepository(db *gorm.DB) *courseRepository{
   return &courseRepository{db: db}
}