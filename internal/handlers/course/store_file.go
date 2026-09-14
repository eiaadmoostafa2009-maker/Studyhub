package course

import (
	"net/http"
	dto "studyhub/internal/dto/course"

	"github.com/gin-gonic/gin"
)

func (h *courseHandler) StoreFile(c *gin.Context) {
	var req dto.UploadFileRequest
    ctx := c.Request.Context()

	if err := c.ShouldBind(&req); err != nil{
		c.HTML(http.StatusBadRequest, "error.html", gin.H{"error": err.Error()})
		return
	}
	header, err := c.FormFile("file")
	if err != nil {
		c.HTML(http.StatusBadRequest, "error.html", gin.H{
			"error": "file is required",
		})
		return
	}

	file, err := header.Open()
	if err != nil {
		c.HTML(500, "error.html", gin.H{
			"error": "could not open file",
		})
		return
	}
	defer file.Close()

	statusCode, result, err := h.service.StoreFile(ctx, file, req)

	if err != nil {
		c.HTML(statusCode, "error.html", gin.H{
			"error": err.Error(),
		})
		return
	}

	c.HTML(statusCode, "file_uploaded.html", gin.H{
		"file": result,
	})
}