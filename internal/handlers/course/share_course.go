package course

import (
	"net/http"
	dto "studyhub/internal/dto/course"

	"github.com/gin-gonic/gin"
)

func(h *courseHandler) ShareCourse(c *gin.Context){
	var(
		ctx =c.Request.Context()
		req dto.ShareCourseRequest
	)

	//put data in json
    err := c.ShouldBind(&req)
	if err != nil{
		c.HTML(http.StatusBadRequest, "error.html", gin.H{"error": err.Error()})
		return
	}
	
	if err := h.validator.Struct(&req); err != nil{
		c.HTML(http.StatusBadRequest, "error.html", gin.H{"error": err.Error()})
		return
	}
	
	//get user id
	userID :=c.GetInt("user_id")
	statusCode, err :=h.service.ShareCourse(ctx, &req, userID)
	if err != nil{
		c.HTML(statusCode, "error.html", gin.H{"error": err.Error()})
		return
	}
	//return response
	c.HTML(statusCode, "message.html", gin.H{
		"message": "uploaded successfully",
	})
}