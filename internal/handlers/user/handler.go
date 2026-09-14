package user

import (
	"net/http"
	"studyhub/internal/middleware"
	"studyhub/internal/models"
	"studyhub/internal/service/user"
    userRepo"studyhub/internal/repository/user"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type userHandler struct {
	api *gin.Engine
	validate *validator.Validate
	service user.UserService
}

func NewUserHandler(api *gin.Engine, validate *validator.Validate, service user.UserService) *userHandler {
	return &userHandler{api: api, validate: validate, service: service}
}

func (h *userHandler) RegisterRoutes(secretkey string, userRepo userRepo.UserRepository) {
	router := h.api.Group("/user")

	router.POST("/register", h.Register)
	router.GET("/register", func (c *gin.Context){
		c.HTML(http.StatusOK, "register.html", nil)
	})
	router.POST("/login", h.Login)
	router.GET("/login", func (c *gin.Context){
		c.HTML(http.StatusOK, "login.html", nil)
	})
	router.POST("/refresh", h.RefreshToken)

	//authenticated routes
	router.Use(middleware.AuthMiddleWare(secretkey))
	router.GET("/home", h.HomePage)
	router.POST("/:id/delete", middleware.RequireRoles(userRepo, models.AdminRole), h.DeleteUser)
	router.GET("/:id/delete", func(c *gin.Context){
		c.HTML(http.StatusOK, "my-profile.html", nil)
	})
	router.GET("/profile", h.Profile)
	router.POST("/profile", h.Profile)
}