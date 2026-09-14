package user

import (
	"net/http"
	dto "studyhub/internal/dto/user"

	"github.com/gin-gonic/gin"
)

func (h *userHandler) Register(c *gin.Context) {
	var (
		req dto.RegisterRequest
		ctx = c.Request.Context()
	)
	if err := c.ShouldBind(&req); err != nil {
		c.HTML(400, "register.html", gin.H{"message": "invalid credentials"})
		return
	}
	
	if err := h.validate.Struct(req); err != nil {
		c.HTML(400, "register.html", gin.H{"message": "invalid credentials"})
		return
	}

    statusCode, err := h.service.Register(ctx, &req)
	if err != nil {
		c.HTML(statusCode, "register.html", gin.H{"message": "invalid credentials"})
		return
	}
	
	c.Header("HX-Redirect", "/user/login")
	c.Status(http.StatusCreated)
}