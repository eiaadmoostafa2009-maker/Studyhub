package post

import(
	"github.com/gin-gonic/gin"
	dto "studyhub/internal/dto/post"
	"net/http"
)

func (h *postHandler) CreatePost(c *gin.Context) {
	var (
		req dto.CreatePostRequest
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
	statusCode, err := h.service.CreatePost(ctx, &req,userID)
	if err != nil {
		c.HTML(statusCode, "error.html", gin.H{"error": err.Error()})
		return
	}
	c.HTML(statusCode, "post_success.html", gin.H{
		"message": "Post created successfully",
	})
}