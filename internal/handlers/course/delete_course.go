package course

import (
	"net/http"
	"strconv"
	dto "studyhub/internal/dto/course"

	"github.com/gin-gonic/gin"
)

func(h *courseHandler) DeleteCourse(c *gin.Context){
	//define context var
	var ctx = c.Request.Context()

	userID :=c.GetInt("user_id")
	courseIDstr :=c.Param("courseID")
	courseID,err := strconv.Atoi(courseIDstr)
	if err != nil{
		c.HTML(http.StatusBadRequest, "error.html", gin.H{"error": err.Error()})
		return
	}

	statusCode,err :=h.service.DeleteCourse(ctx, courseID, userID)
	if err != nil{
		c.HTML(statusCode, "error.html", gin.H{"error": err.Error()})
		return
	}

	c.HTML(statusCode, "message.html", dto.DeleteCourseResponse{
		Message: "The course is deleted",
	})
}