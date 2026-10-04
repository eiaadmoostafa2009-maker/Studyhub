package course

import (
	"log"
	"net/http"
	dto "studyhub/internal/dto/course"

	"github.com/gin-gonic/gin"
)
func (h *courseHandler) SubmitCourse(c *gin.Context) {
    ctx := c.Request.Context()

    var req dto.SubmitCourseRequest

    if err := c.ShouldBind(&req); err != nil {
        c.HTML(http.StatusBadRequest, "error.html", gin.H{
            "error": err.Error(),
        })
        log.Print(err)
        return
    }

    if err := h.validator.Struct(&req); err != nil {
        c.HTML(http.StatusBadRequest, "error.html", gin.H{
            "error": err.Error(),
        })
        log.Print(err)
        return
    }

    file, header, err := c.Request.FormFile("video")
    if err != nil {
        c.HTML(http.StatusBadRequest, "error.html", gin.H{
           "error": "video is required",
        })
        log.Print(err)
        return
    }

    defer file.Close()

    userID := c.GetInt("user_id")
   

    statuscode, err := h.service.SubmitCourse(
      ctx,
      req,
      file,
      header.Filename,
      userID,
    )

    if err != nil {
        c.HTML(statuscode, "error.html", gin.H{
            "error": err.Error(),
        })
        log.Print(err)
        return
    }

    c.HTML(statuscode, "message.html", gin.H{
        "message": "Course submitted for admin review",
    })
}