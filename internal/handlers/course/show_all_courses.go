package course

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func(h *courseHandler) ShowAllCourses(c *gin.Context){
	ctx :=c.Request.Context()

	statusCode,courses,err :=h.service.ShowAllCourses(ctx)
	if err != nil{
		c.HTML(http.StatusBadRequest, "error.html", gin.H{"error": err.Error()})
		return
	}
    
	c.HTML(statusCode, "courses.html", gin.H{
		"courses": courses,
	})
}