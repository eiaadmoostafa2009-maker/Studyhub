package course

import (
	"log"
	"net/http"
	"strconv"
	dto "studyhub/internal/dto/course"

	"github.com/gin-gonic/gin"
)

func(h *courseHandler) PublishCourse(c *gin.Context){
	var(
		ctx = c.Request.Context()
		req dto.PublishCourseRequest
	)

	if err := c.ShouldBind(&req); err != nil{
		c.HTML(http.StatusInternalServerError, "error.html", gin.H{
			"error": err.Error(),
		})
		log.Print(err)
		return
	}

    if err := h.validator.Struct(&req); err != nil{
		c.HTML(http.StatusBadRequest, "error.html", gin.H{
			"error": err.Error(),
		})
		log.Print(err)
		return
	}

	courseIDstr := c.Param("courseID")
	courseID, err := strconv.Atoi(courseIDstr)
	if err != nil {
		c.HTML(http.StatusBadRequest, "error.html", gin.H{
			"error": "invalid course ID",
		})
		log.Print(err)
		return
	}

	statuscode, err := h.service.PublishCourse(ctx, courseID, req)
	if err != nil {
		c.HTML(statuscode, "error.html", gin.H{
			"error": err.Error(),
		})
		log.Print(err)
		return
	}

	c.HTML(statuscode, "message.html", gin.H{
		"message": "Course published successfully",
	})
}