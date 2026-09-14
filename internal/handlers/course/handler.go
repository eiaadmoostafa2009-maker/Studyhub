package course

import (
	"studyhub/internal/middleware"
	"studyhub/internal/models"
	"studyhub/internal/service/course"
    userRepo"studyhub/internal/repository/user"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type courseHandler struct{
	api *gin.Engine
	validator *validator.Validate
	service course.CourseService
}

func NewCourseHandler(api *gin.Engine, validate *validator.Validate, service course.CourseService) courseHandler{
	return courseHandler{
		api: api,
		validator: validate,
		service: service,
	}
}

func (h *courseHandler) RouteList(secretKey string, userRepo userRepo.UserRepository){
	routes :=h.api.Group("course")
	routes.Use(middleware.AuthMiddleWare(secretKey))
	routes.POST("/",middleware.RequireRoles(userRepo, models.TeacherRole), h.ShareCourse)
	routes.POST("/upload",middleware.RequireRoles(userRepo, models.TeacherRole), h.StoreFile)
	routes.POST("/:courseID/update",middleware.RequireRoles(userRepo, models.AdminRole, models.TeacherRole), h.UpdateCourse)
	routes.POST("/:courseID/delete",middleware.RequireRoles(userRepo, models.AdminRole, models.TeacherRole), h.DeleteCourse)
	routes.GET("/",middleware.RequireRoles(userRepo, models.StudentRole, models.TeacherRole, models.AdminRole), h.ShowAllCourses)
}