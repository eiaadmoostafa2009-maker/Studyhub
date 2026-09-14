package user

import (
	dto "studyhub/internal/dto/user"
    "net/http"
	"github.com/gin-gonic/gin"
)

func (h *userHandler) Login(c *gin.Context) {
	var (
		req dto.LoginRequest
		ctx = c.Request.Context()
	)
	if err := c.ShouldBind(&req); err != nil {
		c.HTML(400, "login.html",gin.H{"message": "invalid credentials"})
		return
	}
	if err := h.validate.Struct(req); err != nil {
		c.HTML(400, "login.html", gin.H{"message": "invalid credentials"})
		return
	}
	token, refreshToken, statusCode, err := h.service.Login(ctx, &req)
	if err != nil {
		c.HTML(statusCode, "login.html", gin.H{"message": "invalid credentials"})
		return
	}
	
	c.SetCookie("token", token, 15*60, "/", "", false, true)
	c.SetCookie("refresh_token", refreshToken, 7*24*60*60, "/", "", false, true)
	switch statusCode {
	case 200:
		if c.GetHeader("HX-Request") == "true" {
			c.Header("HX-Redirect", "/user/home")
			c.Status(http.StatusOK)
			return
		}
		c.Redirect(http.StatusSeeOther, "/user/home")
	case 404:
		c.HTML(404, "login.html", gin.H{"message": "invalid credentials"})
	case 401:
		c.HTML(401, "login.html", gin.H{"message": "invalid credentials"})
	case 500:
		c.HTML(500, "login.html", gin.H{"message": "internal server error"})
	}
}