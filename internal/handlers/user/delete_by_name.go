package user

import (
	"net/http"
	dto "studyhub/internal/dto/user"
	"studyhub/internal/models"

	"github.com/gin-gonic/gin"
)

func (h *userHandler) DeleteByName(c *gin.Context) {
	ctx := c.Request.Context()
    var req dto.DeleteByName
	
	if err := c.ShouldBind(&req); err != nil{
		c.HTML(http.StatusBadRequest, "error.html", gin.H{"error": err.Error()})
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
		statuscode, err := h.service.DeleteByName(ctx, req)
		if err != nil {
			c.HTML(statuscode, "message.html", gin.H{"Type": "error", "Message": err.Error()})
			return
		}
		c.HTML(statuscode, "message.html", gin.H{"Type": "success", "Message": "User deleted successfully"})
		return
	}
	c.HTML(http.StatusForbidden, "error.html", gin.H{"error": "forbidden"})
}