package post

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *postHandler) DeletePost(c *gin.Context) {
	var (
		ctx = c.Request.Context()
	)

	userID := c.GetInt("user_id")
	postIDstr := c.Param("post_id")
	postID, err := strconv.Atoi(postIDstr)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", gin.H{"error": err.Error()})
		return
	}

	statusCode, err := h.service.DeletePost(ctx, userID, postID)
	if err != nil {
		c.HTML(statusCode, "error.html", gin.H{"error": err.Error()})
		return
	}
	c.HTML(statusCode, "post_success.html", gin.H{
		"message": "the post have been deleted successfully",
	})
}
