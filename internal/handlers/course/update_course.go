package course

import (
	"net/http"
	"strconv"
	dto "studyhub/internal/dto/course"

	"github.com/gin-gonic/gin"
)

func(h *courseHandler) UpdateCourse(c *gin.Context){
	var(
	  ctx =c.Request.Context()
      req dto.UpdateCourseRequest
	)
    
	err :=c.ShouldBind(&req)
	if err != nil{
       c.HTML(http.StatusBadRequest, "error.html", gin.H{"error": err.Error()})
	   return
	}

	if err := h.validator.Struct(&req); err != nil{
		c.HTML(http.StatusBadRequest, "error.html", gin.H{"error": err.Error()})
		return
	}

	userID :=c.GetInt("user_id")
	idstr :=c.Param("courseID")
    id,_ :=strconv.Atoi(idstr)
	statusCode,err :=h.service.UpdateCourse(ctx, req, id, userID)
    if err != nil{
		c.HTML(statusCode, "error.html", gin.H{"error": err.Error()})
		return
	}

	c.HTML(statusCode, "message.html", gin.H{
		"message": "updated successfully",
	})
}