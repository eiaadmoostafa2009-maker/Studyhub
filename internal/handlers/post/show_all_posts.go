package post

import (
	"github.com/gin-gonic/gin"
)

 func (h *postHandler) ShowAllPosts(c *gin.Context){
	ctx :=c.Request.Context()
	userID := c.GetInt("user_id")

	//get response 
	statusCode, posts, err := h.service.ShowAllPosts(ctx,userID)
	if err != nil{
		c.JSON(statusCode, gin.H{
            "message": err.Error(),
		})
		return
	}
    
	c.HTML(statusCode, "post_list.html", gin.H{
		"posts": posts,
	})
 }