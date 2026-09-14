package post

import(
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *postHandler) JoinPost(c *gin.Context){
	
	ctx := c.Request.Context()

	userID := c.GetInt("user_id")
	postIDstr := c.Param("post_id")
	postID, err :=strconv.Atoi(postIDstr)
	if err != nil {
		c.HTML(http.StatusBadRequest, "error.html", gin.H{"error": err.Error()})
		return
	}

    //get params
	

	statusCode, err := h.service.JoinPost(ctx,userID,postID)
	if err != nil {
		c.HTML(statusCode, "error.html", gin.H{"error": err.Error()})
		return
	}
	c.HTML(statusCode, "post_success.html", gin.H{
		"message": "Joined successfully",
	})
}