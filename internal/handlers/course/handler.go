package course

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"studyhub/internal/middleware"
	"studyhub/internal/models"
	userRepo "studyhub/internal/repository/user"
	"studyhub/internal/service/course"
)

type courseHandler struct {
	api       *gin.Engine
	validator *validator.Validate
	service   course.CourseService
}

func NewCourseHandler(api *gin.Engine, validate *validator.Validate, service course.CourseService) courseHandler {
	return courseHandler{
		api:       api,
		validator: validate,
		service:   service,
	}
}

func (h *courseHandler) RouteList(secretKey string, userRepo userRepo.UserRepository) {
	routes := h.api.Group("/course")
	routes.Use(middleware.AuthMiddleWare(secretKey))
    
	//publish course
	routes.POST("/:courseID/publish", middleware.RequireRoles(userRepo, models.AdminRole), h.PublishCourse)
    //submit request
	routes.POST("/upload", middleware.RequireRoles(userRepo, models.TeacherRole), h.SubmitCourse)
    //updating request
	routes.POST("/:courseID/update", middleware.RequireRoles(userRepo, models.AdminRole, models.TeacherRole), h.UpdateCourse)
    //deleting request
	routes.POST("/:courseID/delete", middleware.RequireRoles(userRepo, models.AdminRole, models.TeacherRole), h.DeleteCourse)
    //showing all courses
	routes.GET("/", middleware.RequireRoles(userRepo, models.StudentRole, models.TeacherRole, models.AdminRole), h.ShowAllCourses)
}
