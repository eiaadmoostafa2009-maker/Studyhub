package post

import (
	"studyhub/internal/middleware"
	"studyhub/internal/models"
	"studyhub/internal/repository/user"
	"studyhub/internal/service/post"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type postHandler struct {
	api      *gin.Engine
	validate *validator.Validate
	service  post.PostService
}

func NewPostHandler(api *gin.Engine, validate *validator.Validate, service post.PostService) postHandler {
	return postHandler{
		api:      api,
		validate: validate,
		service:  service,
	}
}

func (h *postHandler) RouteList(secretKey string, userRepo user.UserRepository) {
	routes := h.api.Group("/posts")
	routes.Use(middleware.AuthMiddleWare(secretKey))
	routes.POST("/", middleware.RequireRoles(userRepo, models.StudentRole), h.CreatePost)
	routes.PUT("/:post_id/update", middleware.RequireRoles(userRepo, models.StudentRole), h.UpdatePost)
	routes.DELETE("/:post_id/delete", middleware.RequireRoles(userRepo, models.StudentRole, models.AdminRole), h.DeletePost)
	routes.POST("/join/:post_id", middleware.RequireRoles(userRepo, models.StudentRole), h.JoinPost)
    routes.GET("/", middleware.RequireRoles(userRepo, models.StudentRole, models.AdminRole),h.ShowAllPosts)
}
