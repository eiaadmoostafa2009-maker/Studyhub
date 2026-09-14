package user

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *userHandler) Profile(c *gin.Context) {
	userIDstr, exists := c.Get("user_id")
	if !exists {
		c.HTML(http.StatusUnauthorized, "error.html", gin.H{"error": "you are not logged in"})
		return
	}

	userID, ok := userIDstr.(int)
	if !ok {
		c.HTML(http.StatusInternalServerError, "error.html", gin.H{"error": "internal error"})
		return
	}

	statusCode, user, err := h.service.GetUserByID(c.Request.Context(), userID)
	if err != nil {
		c.HTML(statusCode, "error.html", gin.H{"error": err.Error()})
		return
	}

	c.HTML(statusCode, "my-profile.html", gin.H{
		"user": user,
	})
}
