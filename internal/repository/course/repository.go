package course

import (
	"context"
	"studyhub/internal/models"

	"gorm.io/gorm"
)

type CourseRepository interface{
    ShareCourse(ctx context.Context, course models.Course)(error)
	StoreFile(ctx context.Context, file models.File)(int, error)
	GetCourseByID(ctx context.Context, id int)(*models.Course, error)
	UpdateCourse(ctx context.Context, id int, course *models.Course) error
	DeleteCourse(ctx context.Context, id int, userID int)error
	ShowAllCourses(ctx context.Context)([]models.CourseWithFile,error)
	GetUserByID(ctx context.Context, userID int) (string, error)
	GetFileByHash (ctx context.Context,	hash string,) (*models.File, error)
}

type courseRepository struct{
	db *gorm.DB
}

func NewCourseRepository(db *gorm.DB) *courseRepository{
   return &courseRepository{db: db}
}