package user

import (
	"net/http"
	"studyhub/internal/models"

	"github.com/gin-gonic/gin"
)

func (h *userHandler) HomePage(c *gin.Context) {
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
		c.HTML(statusCode, "message.html", gin.H{"Type": "error", "Message": err.Error()})
		return
	}

	switch user.Role {
	case models.AdminRole:
		c.HTML(statusCode, "admin-home.html", gin.H{"user": user})
	case models.StudentRole:
		c.HTML(statusCode, "student-home.html", gin.H{"user": user})
	case models.TeacherRole:
		c.HTML(statusCode, "teacher-home.html", gin.H{"user": user})
	default:
		c.HTML(http.StatusInternalServerError, "message.html", gin.H{"Type": "error", "Message": "Internal error"})
		return
	}
}
