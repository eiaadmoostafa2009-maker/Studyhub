package user

import (
	"net/http"
	"studyhub/internal/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *userHandler) DeleteUser(c *gin.Context) {
	ctx := c.Request.Context()

	idstr := c.Param("id")
	id, err := strconv.Atoi(idstr)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", gin.H{"error": err.Error()})
		return
	}

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

	if user.Role == models.AdminRole {
		statuscode, err := h.service.DeleteUser(ctx, id)
		if err != nil {
			c.HTML(statuscode, "message.html", gin.H{"Type": "error", "Message": err.Error()})
			return
		}
		c.HTML(statuscode, "message.html", gin.H{"Type": "success", "Message": "User deleted successfully"})
		return
	}
	c.HTML(http.StatusForbidden, "error.html", gin.H{"error": "forbidden"})
}
