package post

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	dto "studyhub/internal/dto/post"
)

func (h *postHandler) UpdatePost(c *gin.Context) {
	var (
		req dto.UpdatePostRequest
		ctx = c.Request.Context()
	)
	if err := c.ShouldBind(&req); err != nil {
		c.HTML(http.StatusBadRequest, "error.html", gin.H{"error": err.Error()})
		return
	}

	if err := h.validate.Struct(&req); err != nil{
		c.HTML(http.StatusBadRequest, "error.html", gin.H{"error": err.Error()})
		return
	}

	userID := c.GetInt("user_id")
	postIDstr := c.Param("post_id")
	postID, err := strconv.Atoi(postIDstr)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", gin.H{"error": err.Error()})
		return
	}

	statusCode, err := h.service.UpdatePost(ctx, &req, postID, userID)
	if err != nil {
		c.HTML(statusCode, "error.html", gin.H{"error": err.Error()})
		return
	}
	c.HTML(statusCode, "post_success.html", gin.H{
		"message": "the post have been updated successfully",
	})
}

