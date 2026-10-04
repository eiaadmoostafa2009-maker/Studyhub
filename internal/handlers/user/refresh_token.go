package user

import (
	"net/http"
	dto "studyhub/internal/dto/user"

	"github.com/gin-gonic/gin"
)

func (h *userHandler) RefreshToken(c *gin.Context) {
	ctx := c.Request.Context()
	var req dto.RefreshTokenRequest

	if err := c.ShouldBind(&req); err !=nil{
		c.HTML(http.StatusBadRequest, "error.html", gin.H{"error": err.Error()})
		return
	}
	
	if err := h.validate.Struct(req); err != nil {
		c.HTML(400, "register.html", gin.H{"message": "invalid credentials"})
		return
	}

	userID := c.GetInt("user_id")
	refreshToken, err := c.Cookie("refresh_token")
	if err != nil {
		c.HTML(http.StatusUnauthorized, "error.html", gin.H{"error": err.Error()})
		return
	}

	refreshToken, token, statusCode, err := h.service.RefreshToken(ctx, &req,int(userID), refreshToken)
	if err != nil {
		c.HTML(statusCode, "error.html", gin.H{"error": err.Error()})
		return
	}
	
	c.SetCookie(
       "token",
       token,
       15*60, // 15 minutes
       "/",
       "",
       false, // true in HTTPS production
       true,  // HttpOnly
    )

    c.SetCookie(
      "refresh_token",
      refreshToken,
      7*24*60*60,
      "/user/refresh",
      "",
      false,
      true,
    )

   c.Header("HX-Redirect", "/")
   c.Status(http.StatusNoContent)
}